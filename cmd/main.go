package main

import (
	"golang-templ/views"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := views.Hello("john")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(component).ServeHTTP(w, r)
	})

	// Added error handling for ListenAndServe
	if err := http.ListenAndServe(":8000", nil); err != nil {
		// Log the error and exit
		panic(err)
	}
}
