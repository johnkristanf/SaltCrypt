package database


import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)


func MigrateUsersTable() error {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	// SQL query to create the users table
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		salt TEXT NOT NULL,
		pepper TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
	)`

	// Execute the query
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	fmt.Println("Migration executed successfully! Users table created.")
	return nil
}

func ConnectDB() *pgx.Conn {
	
	dsn := os.Getenv("DSN")

	var err error
	dbConn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to the database")

	return dbConn
}
