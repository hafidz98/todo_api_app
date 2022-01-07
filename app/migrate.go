package app

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hafidz98/todo_api_app/helper"
	"github.com/joho/godotenv"
)

func NewMigrate() {
	log.Println("Database migrations start")

	err := godotenv.Load(".env")
	helper.PanicIfError(err)

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	// := "root@tcp(localhost:3306)/todos_api_db"
	dsn := (username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?multiStatements=true")
	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	//db, _ := sql.Open("mysql", "user:password@tcp(host:port)/dbname?multiStatements=true")
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	helper.PanicIfError(err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		dbname,
		driver,
	)
	helper.PanicIfError(err)

	m.Up()
}
