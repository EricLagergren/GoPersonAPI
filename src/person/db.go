package person

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Database variable useable from anywhere within our application
var db *sql.DB

// Runs when the Go file is initialized
func init() {
	// Prepare our database connection string
	d := fmt.Sprintf("%s:%s@tcp4(%s:%s)/%s",
		os.Getenv("MYSQLUSER"),
		os.Getenv("MYSQLPASS"),
		os.Getenv("MYSQLHOST"),
		os.Getenv("MYSQLPORT"),
		os.Getenv("MYSQLDB"),
	)

	var err error
	db, err = sql.Open("mysql", d)
	if err != nil {
		panic(err)
	}

	// Ping does not open a connection, will only validate DSN data
	// Something went wrong with our DSN data sent to the MySQL server
	if err := database.Ping(); err != nil {
		panic(err)
	}

	log.Print("Database connected successfully!")
}
