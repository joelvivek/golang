package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "test title", Desc: " test description", Content: "content of article"},
	}

	fmt.Println("Endpoint hit : all article")
	json.NewEncoder(w).Encode(articles)
}

func PostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST method Worked")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", PostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func main() {
	handleRequest()
}
