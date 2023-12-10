package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type MySQL struct {
	DB *sql.DB
}

func (m *MySQL) Connect() error {
	dbHost := os.Getenv("MYSQL_DB_HOST")
	dbPort := os.Getenv("MYSQL_DB_PORT")
	dbName := os.Getenv("MYSQL_DB_DATABASE")
	dbUser := os.Getenv("MYSQL_DB_USERNAME")
	dbPass := os.Getenv("MYSQL_DB_PASSWORD")

	// Create MySQL DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// Ping database to verify connection
	err = db.Ping()
	if err != nil {
		return err
	}

	// Assign the opened connection to the struct field
	m.DB = db

	fmt.Println("Connected to the MySQL database!")
	return nil
}
