package main

import (
	"go-book-mysql-abs/pkg/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.Connect()
}
