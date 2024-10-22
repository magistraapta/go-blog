package main

import (
	"fmt"
	"golang-templ/cmd/database"
	"log"
	"net/http"
)

func main() {

	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}
	log.Println("Success connect to DB")
	// Added error handling for ListenAndServe
	fmt.Print("Running server on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		// Log the error and exit
		panic(err)
	}
}
