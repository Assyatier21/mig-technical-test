package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), "mig-attendance", os.Getenv("DB_SCHEMA"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Println("[Database] can't connect to database: ", err.Error())
		return nil
	}

	log.Println("[Database] successfully connected")
	return db
}
