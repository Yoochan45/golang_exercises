package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Connection struct {
	Host     string
	Dbname   string
	Username string
	Password string
	Port     string
}

func connectionParser() *Connection {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when load env", err)
	}

	host := os.Getenv("DATABASE_HOST")
	dbname := os.Getenv("DATABASE_DBNAME")
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	port := os.Getenv("DATABASE_PORT")

	return &Connection{
		Host:     host,
		Dbname:   dbname,
		Username: username,
		Password: password,
		Port:     port,
	}
}
