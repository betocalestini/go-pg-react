package api

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (string, *sql.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbDatabase := os.Getenv("DB_DATABASE")
	DbDriver := os.Getenv("DB_DRIVER")
	ServerAddress := os.Getenv("SERVER_ADDRESS")

	conn, err := sql.Open(DbDriver, "postgresql://"+DbUser+":"+DbPassword+"@"+DbHost+":"+DbPort+"/"+DbDatabase+"?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	return ServerAddress, conn
}
