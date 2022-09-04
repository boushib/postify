package config

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

type Connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func InitDBConnection() {
	envMap, err := godotenv.Read(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	connection := Connection{
		Host:     envMap["DB_HOST"],
		Port:     envMap["DB_PORT"],
		User:     envMap["DB_USER"],
		Password: envMap["DB_PASSWORD"],
		DBName:   envMap["DB_NAME"],
	}

	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", connection.Host, connection.Port, connection.User, connection.Password, connection.DBName))

	if err != nil {
		fmt.Println("Error connecting to database")
		return
	}

	fmt.Println("Successfully connected to database!")

	err = DB.Ping()

	if err != nil {
		fmt.Println("Error pinging database")
		return
	}

	fmt.Println("Successfully pinged database!")
}
