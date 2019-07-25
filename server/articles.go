package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Article Example Documentation
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func initArticles() {
	Articles = []Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	switch vars["id"] {
	case "1":
		json.NewEncoder(w).Encode(Articles[0])
	case "2":
		json.NewEncoder(w).Encode(Articles[1])
	default:
		fmt.Fprintf(w, "bugger off")
	}

	return
}
