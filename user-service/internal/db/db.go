package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var PI PostgresInstance

type PostgresInstance struct {
	DB *sql.DB
}

func PostgreSQLConnection() {
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
	uri := fmt.Sprintf("postgresql://%s:%s@user-postgres:5432/%s?sslmode=disable",
		os.Getenv("USER_DB_USERNAME"),
		os.Getenv("USER_DB_PASSWORD"),
		os.Getenv("USER_DB_NAME"),
	)

	// Define database connection for PostgreSQL.
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}

	// Set database connection settings.
	db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
	db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close()
		log.Fatal(err)
	}

	// Create users table if does not exist
	if err := createUsersTable(db); err != nil {
		log.Fatal(err)
	}

	// Set Postgres instance
	PI = PostgresInstance{
		DB: db,
	}
}

func createUsersTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id serial primary key,
		email varchar(50),
		username varchar(50),
		display_name varchar(50),
		pwd varchar(100),
		is_admin boolean,
		created_at timestamp,
		updated_at timestamp
	)`

	_, err := db.Exec(query)
	return err
}
