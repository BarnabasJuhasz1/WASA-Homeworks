/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"net/http"
	"sapienza/wasatext/service/api/reqcontext"
	"sapienza/wasatext/service/api/util"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetUser(userID int) (util.User, error)
	GetUserFromName(username string) (util.User, error)
	InsertUser(newUser util.User) (int, error)
	UpdateUser(newUser util.User, oldUsername string) error
	GetLoggedInUser(w http.ResponseWriter, ctx reqcontext.RequestContext) util.User

	AddConversationIDToUser(userID int, addToConversationId int) error
	RemoveConversationIDFromUser(userID int, removeFromConversationId int) error

	GetConversation(id int) (util.Conversation, error)
	InsertConversation(newConversation util.Conversation) (int, error)
	GetMyConversationIDs(userID int) ([]int, error)
	GetMyConversations(userID int) ([]util.Conversation, error)
	UpdateConversation(id int, conversation util.Conversation) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='users'").Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		// fmt.Println("users Table Created")
		sqlStmt := `CREATE TABLE users (
					id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
					username TEXT,
					profile_picture TEXT,
					conversations JSON);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, errors.New("error creating database structure: TABLE users")

		}
	}
	err2 := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='conversations'").Scan(&tableName)
	if errors.Is(err2, sql.ErrNoRows) {
		// fmt.Println("conversations Table Created")
		sqlStmt := `CREATE TABLE conversations (
					id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
					type TEXT NOT NULL,
					group_name TEXT,
					group_picture BLOB, 
					participants JSON,
					Messages JSON);`
		_, err2 = db.Exec(sqlStmt)
		if err2 != nil {
			return nil, errors.New("error creating database structure: TABLE conversations")

		}
	}
	// err3 := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='user_to_conversations'").Scan(&tableName)
	// if errors.Is(err3, sql.ErrNoRows) {
	// 	fmt.Println("user_to_conversations Table Created")
	// 	sqlStmt := `CREATE TABLE user_to_conversations (
	// 				username TEXT NOT NULL PRIMARY KEY,
	// 				conversations JSON);`
	// 	_, err3 = db.Exec(sqlStmt)
	// 	if err3 != nil {
	// 		return nil, fmt.Errorf("error creating database structure: %w", err3)

	// 	}
	// }

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
