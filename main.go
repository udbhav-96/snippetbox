package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func getHome(w http.ResponseWriter, r *http.Request){
	fmt.Println("got request - Home")
	fmt.Fprintf(w, "Hello From Snippetbox")
}

func snippetView(w http.ResponseWriter, r *http.Request){
	fmt.Println("got request - Snippet View")
	fmt.Fprintf(w, "Hello From Snippetbox - View")
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
	fmt.Println("got request - Snippet Create")
	fmt.Fprintf(w, "Hello From Snippetbox - Create")
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", getHome)
	r.HandleFunc("/view", snippetView)
	r.HandleFunc("/create", snippetCreate)

	err := http.ListenAndServe(":4000", r)
	log.Fatal(err)

}