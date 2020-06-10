package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

//Article is ..
type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}
//Articles is ...
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request)  {
	articles := Articles{
		Article{Title:"Test Title", Desc : "Test Desc", Content: "Test Content"},
	}
	fmt.Println("***************")
	fmt.Println(r);
	fmt.Println("***************")
	fmt.Println("Getting all articles from server")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	handleRequest();
}
