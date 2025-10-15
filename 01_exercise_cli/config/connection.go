package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	con := connectionParser()

	fmt.Println(con)

	dsn := fmt.Sprintf("host=%s dbname=%s username=%s password=%s port=%s",
		con.Host, con.Dbname, con.Username, con.Password, con.Port)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error:", err)
	}

	return db
}
