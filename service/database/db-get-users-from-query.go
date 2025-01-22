package database

// GetUsersBySearch returns the userIDs in the database matching the given search query
func (db *appdbimpl) GetUsersFromQuery(searchQuery string) ([]int, error) {

	var userIDs []int

	rows, err := db.c.Query("SELECT id FROM users WHERE username LIKE (?)", "%"+searchQuery+"%")
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
