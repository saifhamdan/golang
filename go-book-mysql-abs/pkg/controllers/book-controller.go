package controllers

import (
	"context"
	"database/sql"
	"go-book-mysql-abs/pkg/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB // Database connection pool.

// Query the database for the information requested and prints the results.
// If the query fails exit the program with an error.
func GetBookById(ctx context.Context, id int64) {
	config.Connect()

	var name string
	err := pool.QueryRowContext(ctx, "select * from people where p.id = :id;", sql.Named("id", id)).Scan(&name)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}
	log.Println("name=", name)
}
