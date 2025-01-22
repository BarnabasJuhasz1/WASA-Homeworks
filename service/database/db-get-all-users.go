package database

// GetAllUsers returns the userIDs in the database
func (db *appdbimpl) GetAllUsers() ([]int, error) {
	var userIDs []int

	rows, err := db.c.Query("SELECT id FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		rowErr := rows.Scan(&id)
		if rowErr != nil {
			return nil, rowErr
		}
		userIDs = append(userIDs, id)
	}

	err2 := rows.Err()
	if err2 != nil {
		return nil, err2
	}

	return userIDs, nil
}
