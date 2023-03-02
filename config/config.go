package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadENV :nodoc:
func LoadENV() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Sprintf("ERROR: %s", err))
	}

	return err
}

// ENV :nodoc:
func ENV() string {
	return os.Getenv("ENV")
}

// HTTPPort :nodoc:
func HTTPPort() string {
	return os.Getenv("PORT")
}

// DBHost :nodoc:
func DBHost() string {
	return os.Getenv("DB_HOST")
}

// DBUser :nodoc:
func DBUser() string {
	return os.Getenv("DB_USERNAME")
}

// DBPassword :nodoc:
func DBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

// DBName :nodoc:
func DBName() string {
	return os.Getenv("DB_DATABASE")
}

// DBPort :nodoc:
func DBPort() string {
	return os.Getenv("DB_PORT")
}
