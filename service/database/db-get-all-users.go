package database

// GetAllUsers returns the userIDs in the database
func (db *appdbimpl) GetAllUsers() ([]int, error) {

	// initialize all userIDs list
	var userIDs []int

	// query to select all rows
	rows, err := db.c.Query("SELECT id FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// scan the rows one-by-one
	for rows.Next() {
		var id int
		rowErr := rows.Scan(&id)
		if rowErr != nil {
			return nil, rowErr
		}
		// if row is scanned properly, add it to the list
		userIDs = append(userIDs, id)
	}

	err2 := rows.Err()
	if err2 != nil {
		return nil, err2
	}

	// return list
	return userIDs, nil
}
