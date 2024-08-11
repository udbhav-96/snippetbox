package main

import "net/http"

func (app * application) routes() *http.ServeMux{

	mux := http.NewServeMux()
	
	fileServer := http.FileServer(http.Dir("../../ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/view", app.ViewSnippet)
	mux.HandleFunc("/create", app.CreateSnippet)

	return mux

}