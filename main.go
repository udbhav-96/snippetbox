package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"strconv"
)

func getHome(w http.ResponseWriter, r *http.Request){
	fmt.Println("got request - Home")
	fmt.Fprintf(w, "Hello From Snippetbox")
}

func snippetView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w, r)
		return
	}
	fmt.Println("got request - Snippet View")
	fmt.Fprintf(w, "Hello From Snippetbox - View ID: %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405);
		// fmt.Fprintf(w, "Method Not Allowed")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	fmt.Fprintf(w, "Create a new Snippet...")
	fmt.Println("got request - Snippet Create")
	
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", getHome)
	r.HandleFunc("/view", snippetView)
	r.HandleFunc("/create", snippetCreate)

	err := http.ListenAndServe(":4000", r)
	log.Fatal(err)

}