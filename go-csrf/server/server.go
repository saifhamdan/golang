package server

import (
	"csrf/server/middleware"
	"log"
	"net/http"
)

func StartServer(hoostname string, port string) error {
	host := hoostname + ":" + port

	log.Printf("Listening on %s", host)
	handler := middleware.NewHandler()

	http.Handle("/", handler)
	return http.ListenAndServe(host, nil)
}
