package database

import "fmt"

// GetUsersBySearch returns the userIDs in the database matching the given search query
func (db *appdbimpl) GetUsersFromQuery(searchQuery string) ([]int, error) {
	var userIDs []int

	fmt.Println("db asked for QUERY; ", searchQuery)

	rows, err := db.c.Query("SELECT id FROM users WHERE username LIKE ?", "%"+searchQuery+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println("row running QUERY; ")

		var id int
		rowErr := rows.Scan(&id)
		if rowErr != nil {
			return nil, rowErr
		}
		userIDs = append(userIDs, id)
		fmt.Println("row found id: ", id)
	}

	err2 := rows.Err()
	if err2 != nil {
		return nil, err2
	}

	return userIDs, nil
}
