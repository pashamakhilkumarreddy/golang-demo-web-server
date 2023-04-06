package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	handlers "github.com/pashamakhilkumarreddy/golang-web-server/handlers"
)

func main() {
	fmt.Println("Welcome to Golang Web Server")

	r := mux.NewRouter()

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/ping", handlers.HealthCheckHandler).Methods("GET")
	r.HandleFunc("/add", handlers.AddFounderHandler).Methods("POST")

	fmt.Printf("Starting server at port %s \n", port)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
