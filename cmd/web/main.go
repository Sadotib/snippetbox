package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() /*although a DefaultServeMux is provided by Go, we will create our
	own mux to have more control over the routing of our application
	and also for the sake of security*/
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	// routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Server Started on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
