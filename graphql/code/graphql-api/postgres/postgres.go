package postgres

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"graphql-api/logger"

	// postgres driver
	_ "github.com/lib/pq"
)

// Db is our database struct used for interacting with the database
type Db struct {
	*sql.DB
}

// New makes a new database using the connection string and
// returns it, otherwise returns the error
func New(connString string) (*Db, error) {
	logger.Log().Infof("Open connection to database with connection string %s", connString)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Errorf("Failed to open database connection, with error: %+v", err)
		return nil, err
	}

	logger.Log().Infof("Checking db connection...")
	// Check that our connection is good
	if err := db.Ping(); err != nil {
		logger.Log().Errorf("Db connection failed!, with error %+v", err)
		return nil, err
	}
	logger.Log().Infof("Db Successfully connected")

	return &Db{db}, nil
}

// ConnString returns a connection string based on the parameters it's given
// This would normally also contain the password, however we're not using one
func ConnString(host string, port int, user string, dbName string, password string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbName, password,
	)
}

// User shape
type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

// GetUsersByName is called within our user query for graphql
func (d *Db) GetUsersByName(name string) []User {
	logger.Log().Infof("Get user name %s from db", name)
	// Prepare query, takes a name argument, protects from sql injection
	stmt, err := d.Prepare("SELECT * FROM users WHERE name=$1")
	if err != nil {
		logger.Log().Errorf("GetUserByName Preperation Err: %+v", err)
	}

	// Make query with our stmt, passing in name argument
	rows, err := stmt.Query(name)
	if err != nil {
		logger.Log().Errorf("GetUserByName Query Err: %+v", err)
	}

	logger.Log().Infof("User name %s found, constructing results...", name)
	// Create User struct for holding each row's data
	var r User
	// Create slice of Users for our response
	users := []User{}
	// Copy the columns from row into the values pointed at by r (User)
	for rows.Next() {
		if err := rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
			&r.Profession,
			&r.Friendly,
		); err != nil {
			log.Errorf("Error scanning rows: %+v", err)
		}
		users = append(users, r)
	}

	logger.Log().Infof("Db query results is %+v", users)
	return users
}
