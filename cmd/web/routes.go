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

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.Home))
	router.Handler(http.MethodGet, "/view/:id", dynamic.ThenFunc(app.ViewSnippet))
	router.Handler(http.MethodGet, "/create", dynamic.ThenFunc(app.CreateSnippet))
	router.Handler(http.MethodPost, "/create", dynamic.ThenFunc(app.CreateSnippetPost))

	
	// return app.recoverPanic(app.logRequest(secureHeaders(mux))) // old way doing manualy 

	// now doing same thing (middleware ) with justinas/alice package
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}