package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux() /*although a DefaultServeMux is provided by Go, we will create our
	own mux to have more control over the routing of our application
	and also for the sake of security*/

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
