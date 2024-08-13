package main

import (
	"fmt"
	"net/http"
	"strconv"
	"errors"
	"html/template"

	"snippetbox/internal/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		app.notFound(w) // http.NotFound(w, r)

		return 
	}

	files := []string{
		"../../ui/html/base.tmpl.html",
		"../../ui/html/pages/home.tmpl.html",
		"../../ui/html/partials/nav.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil{
		app.serverError(w, err)
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error1",  500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil{
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error2", 500)
		app.serverError(w, err)
	}
}

func (app *application) ViewSnippet(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id<1{
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil{
		if errors.Is(err, models.ErrNoRecord){
			app.notFound(w)
		} else{
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		w.Header().Set("Allow", http.MethodPost)
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
    content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
    expires := 7

    id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
    	app.serverError(w, err)
    	return 
    }

	// w.Write([]byte("Create a new Snippet..."))
	http.Redirect(w, r, fmt.Sprintf("/view?id=%d", id), http.StatusSeeOther)
}