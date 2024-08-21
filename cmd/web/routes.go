package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app * application) routes() http.Handler{

	mux := http.NewServeMux()
	
	fileServer := http.FileServer(http.Dir("../../ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/view", app.ViewSnippet)
	mux.HandleFunc("/create", app.CreateSnippet)

	
	// return app.recoverPanic(app.logRequest(secureHeaders(mux))) // old way doing manualy 

	// now doing same thing (middleware ) with justinas/alice package
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(mux)
}