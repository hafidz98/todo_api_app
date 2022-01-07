package app

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/hafidz98/todo_api_app/helper"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	log.Println("Database connection start")

	err := godotenv.Load(".env")
	helper.PanicIfError(err)

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	// := "root@tcp(localhost:3306)/todos_api_db"
	dsn := (username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname)
	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	log.Println(dsn)
	//log.Println("Database host: port:")

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	if err != nil {
		log.Println("Database connection error")
	}

	log.Println("Database connection established")
	return db
}
