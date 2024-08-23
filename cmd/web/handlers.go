package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox/internal/models"
	"snippetbox/internal/validator"

	"github.com/julienschmidt/httprouter"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request){
	// if r.URL.Path != "/"{
	// 	app.notFound(w) // http.NotFound(w, r)
	// 	return 
	// }

	snippets, err := app.snippets.Latest()
	if err != nil{
		app.serverError(w,err)
		return
	}

	// for _,snippet := range snippets {
	// 	fmt.Fprintf(w, "%+v\n", snippet)
	// }

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, http.StatusOK, "home.tmpl.html", data)

	// files := []string{
	// 	"../../ui/html/base.tmpl.html",
	// 	"../../ui/html/pages/home.tmpl.html",
	// 	"../../ui/html/partials/nav.tmpl.html",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil{
	// 	app.serverError(w, err)
	// 	// app.errorLog.Println(err.Error())
	// 	// http.Error(w, "Internal Server Error1",  500)
	// 	return
	// }

	// data := &templateData{
	// 	Snippets : snippets,
	// }

	// err = ts.ExecuteTemplate(w, "base", data)
	// if err != nil{
	// 	// app.errorLog.Println(err.Error())
	// 	// http.Error(w, "Internal Server Error2", 500)
	// 	app.serverError(w, err)
	// }
}

func (app *application) ViewSnippet(w http.ResponseWriter, r *http.Request){
	// id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// if err != nil || id<1{
	// 	// http.NotFound(w, r)
	// 	app.notFound(w)
	// 	return
	// }

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1{
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil{
		if errors.Is(err, models.ErrNoRecord){
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// files := []string{
	// 	"../../ui/html/base.tmpl.html",
	// 	"../../ui/html/partials/nav.tmpl.html",
	// 	"../../ui/html/pages/view.tmpl.html",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil{
	// 	app.serverError(w, err)
	// 	return
	// }

	// data := &templateData{
	// 	Snippet: snippet,
	// }

	// err = ts.ExecuteTemplate(w, "base", data)
	// if err != nil{
	// 	app.serverError(w, err)
	// }


	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, http.StatusOK, "view.tmpl.html", data)
}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request){
	// w.Write([]byte("Display the form for creating a new snippet..."))
	data := app.newTemplateData(r)

	data.Form = snippetCreateForm{
		Expires: 365,
	}
	app.render(w, http.StatusOK, "create.tmpl.html", data)
}

type snippetCreateForm struct{
	Title 		string	`form:"title"`
	Content		string 	`form:"content"`
	Expires		int 	`form:"expires"`
	validator.Validator `form:"-"`
}

func (app *application) CreateSnippetPost(w http.ResponseWriter, r *http.Request){
	// if r.Method != http.MethodPost{
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }

	// err := r.ParseForm()
	// if err != nil{
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }

    // expires, err := strconv.Atoi(r.PostForm.Get("expires"))
    // if err != nil {
    // 	app.clientError(w, http.StatusBadRequest)
    // 	return
    // }

    // form := snippetCreateForm{
    // 	Title : 	r.PostForm.Get("title"),
    // 	Content : 	r.PostForm.Get("content"),
    // 	Expires: 	expires,
    // }

    

    // // check title
    // if strings.TrimSpace(form.Title) == "" {
    // 	form.FieldErrors["title"] = "This field cannot be blank"
    // } else if utf8.RuneCountInString(form.Title) > 100 {
    // 	form.FieldErrors["title"] = "This field cannot be more than 100 characters long"
    // }

    // // check content
    // if strings.TrimSpace(form.Content) == "" {
    // 	form.FieldErrors["content"] = "This field cannot be blank"
    // }

    // // check expiry value
    // if !(form.Expires == 1 || form.Expires == 7 || form.Expires == 365) {
    // 	form.FieldErrors["expires"] = "This field must equal 1, 7 or 365"
    // }

    // if len(form.FieldErrors) >0 {
    // 	data := app.newTemplateData(r)
    // 	data.Form = form
    // 	app.render(w, http.StatusUnprocessableEntity, "create.tmpl.html", data)
    // 	return
    // }
    var form snippetCreateForm

    err := app.decodePostForm(r, &form)
    if err != nil{
    	app.clientError(w, http.StatusBadRequest)
    	return
    }

    form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
    form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
    form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
    form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal to 1, 7, or 365")

    if !form.Valid(){
    	data := app.newTemplateData(r)
    	data.Form = form
    	app.render(w, http.StatusUnprocessableEntity, "create.tmpl.html", data)
    	return
    }

    id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
    if err != nil {
    	app.serverError(w, err)
    	return 
    }

    app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")

	// w.Write([]byte("Create a new Snippet..."))
	http.Redirect(w, r, fmt.Sprintf("/view/%d", id), http.StatusSeeOther)
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Display a HTML form for signing up a new user...")
}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Create a new user...")
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Display a HTML form for logging in a user...")
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Authenticate and login the user...")
}

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Logout the user...")
}