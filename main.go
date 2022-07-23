package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

//create a new article in articles endpoint
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	fmt.Println("Returning all articles")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test enpoint Endpoint Hit")
}

//homepage endpoint
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

//handle all requests
func handleRequest() {

	// new gorrila mux router instance to access all functions within the library
	myRouter := mux.NewRouter().StrictSlash(true)
	//homepage endpoint
	myRouter.HandleFunc(
		"/", homePage)
	//articles endpoint
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() { handleRequest() }
