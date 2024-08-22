package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app * application) routes() http.Handler{


	router := httprouter.New()
	// mux := http.NewServeMux()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("../../ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.Home)
	router.HandlerFunc(http.MethodGet, "/view/:id", app.ViewSnippet)
	router.HandlerFunc(http.MethodGet, "/create", app.CreateSnippet)
	router.HandlerFunc(http.MethodPost, "/create", app.CreateSnippetPost)

	
	// return app.recoverPanic(app.logRequest(secureHeaders(mux))) // old way doing manualy 

	// now doing same thing (middleware ) with justinas/alice package
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}