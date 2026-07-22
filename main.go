package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() /*although a DefaultServeMux is provided by Go, we will create our
	own mux to have more control over the routing of our application
	and also for the sake of security*/

	//routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Server Started on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Viewing a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST") /*setting the header should be done before
		WriteHeader/Write is called, otherwise no effect*/
		w.WriteHeader(405)

		w.Write([]byte("Method Not Allowed"))

		return
	}
	w.Write([]byte("Creating a new snippet..."))
}
