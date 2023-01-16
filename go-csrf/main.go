package main

import (
	"csrf/db"
	"csrf/server"
	"csrf/server/middleware/myJwt"
	"log"
)

var host = "localhost"
var port = "9000"

func main() {
	db.InitDB()

	jwtErr := myJwt.InitJWT()
	if jwtErr != nil {
		log.Println("Error initializing the JWT!")
		log.Fatal(jwtErr)
	}

	serverErr := server.StartServer(host, port)
	if serverErr != nil {
		log.Println("Error starting server!")
		log.Fatal(serverErr)
	}
}
