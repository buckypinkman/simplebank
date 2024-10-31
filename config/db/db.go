package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Define exported configuration variables
var (
    AppName     string
    DBDriver    string
    DBName      string
    DBPort      string
    DBHost      string
    DBUsername  string
    DBPassword  string
    DBSource    string
)

func init() {
    // Load .env file, if available
    if err := godotenv.Load("../../.env"); err != nil {
        log.Println("No .env file found. Using system environment variables.")
    }

    // Initialize configuration variables with environment variables or default values
    AppName = getEnv("APP_NAME", "Default App Name")
    DBDriver = getEnv("DB_DRIVER", "postgres")
    DBName = getEnv("DB_NAME", "default_db")
    DBPort = getEnv("DB_PORT", "5432")
    DBHost = getEnv("DB_HOST", "locallhost")
    DBUsername = getEnv("DB_USERNAME", "default_user")
    DBPassword = getEnv("DB_PASSWORD", "default_password")
	DBSource = "postgresql://" + DBUsername + ":" + DBPassword + "@" + DBHost + "/" + DBName + "?" + "sslmode=disable"

        fmt.Println(DBSource)
}

// Helper function to retrieve environment variables with fallback to default values
func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
