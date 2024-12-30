package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is the global database connection pool
var DB *sql.DB

// getEnv retrieves environment variables, returns defaultValue if not set
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// InitDB initializes the database connection
func InitDB() error {
	// Retrieve database connection details from environment variables
	dbUser := getEnv("POSTGRES_USER", "root")
	dbPassword := getEnv("POSTGRES_PASSWORD", "1234")
	dbName := getEnv("POSTGRES_DB", "db")
	dbHost := getEnv("POSTGRES_HOST", "db")
	dbPort := getEnv("POSTGRES_PORT", "5432")
	sslMode := getEnv("POSTGRES_SSLMODE", "disable") // Change to "require" in production

	// Build PostgreSQL connection string
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, sslMode,
	)

	log.Printf("Attempting to connect to DB with host=%s, port=%s, user=%s, dbname=%s, sslmode=%s",
		dbHost, dbPort, dbUser, dbName, sslMode)

	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		DB, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("Attempt %d: cannot open DB connection: %v", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}

		// Validate the connection
		if err = DB.Ping(); err != nil {
			log.Printf("Attempt %d: cannot ping DB: %v", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}

		log.Println("Successfully established database connection.")
		return nil
	}

	return fmt.Errorf("failed to connect to DB after %d attempts: %w", maxRetries, err)
}
