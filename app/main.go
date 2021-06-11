package main

import (
	"awesomeProject/app/route"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//creating instance of mux
	routes := route.GetRoutes()
	http.Handle("/", routes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
