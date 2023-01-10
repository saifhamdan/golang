package config

import (
	"database/sql"
	"fmt"
	"go-book-mysql-abs/pkg/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	db, err := sql.Open("mysql", "root:1234@/bookstore")
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	var Book models.Book
	// var names string
	db.QueryRow("SELECT name from books").Scan(&Book.Name)
	fmt.Println(Book.Name)
}
