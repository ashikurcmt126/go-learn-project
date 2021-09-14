package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	Title string `"json:"Title""`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request)  {
	articles := Articles{
		Article{Title: "Test title", Desc: "Test Description", Content: "Hello world"},
	}
	fmt.Println("Endpoint Hit. All articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Test POST endPoint worked")
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "homepage endpoint hit1..")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080",myRouter))
}

func main()  {
	handleRequest()
}