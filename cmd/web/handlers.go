package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Hello from Snippetbox"))
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	sId := r.URL.Query().Get("id")
	iId, err := strconv.Atoi(sId)
	if err != nil || iId < 1 {
		http.Error(w, "Invalid ID", http.StatusNotFound)
		return
	}
	w.Write([]byte("Viewing a specific snippet: " + sId))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost) /*setting the header should be done before
		WriteHeader/Write is called, otherwise no effect*/

		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age=31536000")

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) /*we can use this helper function
		which calls WriteHeader and Write behind the scenes*/
		return
	}
	w.Write([]byte("Creating a new snippet..."))
}
