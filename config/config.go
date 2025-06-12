package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectToDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read from env
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is empty. Check .env file.")
	}

	// Connect to DB
	DB, err = sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	fmt.Println("Connected to PostgreSQL DB successfully!")
}

// package config for local testing , comment out if you test locally

// import (
// 	"fmt"
// 	"log"

// 	"github.com/jmoiron/sqlx"
// 	_ "github.com/lib/pq"
// )

// var DB *sqlx.DB

// func ConnectToDB() {
// 	dbUser := "clinic_user"
// 	dbPassword := "strongpassword"
// 	dbName := "clinic_app"
// 	dbHost := "localhost"
// 	dbPort := "5432"

// 	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
// 		dbUser, dbPassword, dbName, dbHost, dbPort)

// 	var err error
// 	DB, err = sqlx.Open("postgres", dsn)
// 	if err != nil {
// 		log.Fatalf("Error opening database: %v", err)
// 	}

// 	err = DB.Ping()
// 	if err != nil {
// 		log.Fatalf("Error connecting to database: %v", err)
// 	}

// 	fmt.Println("Successfully connected to PostgreSQL!")
// }
