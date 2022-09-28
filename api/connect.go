package api

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (string, *sql.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DbUser := os.Getenv("DbUser")
	DbPassword := os.Getenv("DbPassword")
	DbHost := os.Getenv("DbHost")
	DbPort := os.Getenv("DbPort")
	DbDatabase := os.Getenv("DbDatabase")
	DbDriver := os.Getenv("DbDriver")
	ServerAddress := os.Getenv("ServerAddress")

	conn, err := sql.Open(DbDriver, "postgresql://"+DbUser+":"+DbPassword+"@"+DbHost+":"+DbPort+"/"+DbDatabase+"?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	return ServerAddress, conn
}
