package main

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
	"html/template"

)

func Home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return 
	}

	files := []string{
		"../../ui/html/base.tmpl.html",
		"../../ui/html/pages/home.tmpl.html",
		"../../ui/html/partials/nav.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil{
		log.Println(err.Error())
		http.Error(w, "Internal Server Error1",  500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil{
		log.Println(err.Error())
		http.Error(w, "Internal Server Error2", 500)
	}
}

func ViewSnippet(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id<1{
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying the specific snippet with ID %d...", id)
}

func CreateSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new Snippet..."))
}