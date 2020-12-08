package main

import (
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json: "Title"`
	Desc    string `json: "desc"`
	Content string `json:"content"`
}
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Articles")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage, Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	handleRequests()
}
