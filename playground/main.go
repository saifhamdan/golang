package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func RunWebPortal(addr string) error {
	http.HandleFunc("/", rootHandler)
	return http.ListenAndServe(addr, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Dino web portal %s", r.RemoteAddr)
}

type configuration struct {
	Webserver string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	RunWebPortal(config.Webserver)
}
