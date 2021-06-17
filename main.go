package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//defining struct
type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request)  {
	articles := Articles{
		Article{Title: "First Articule Title", Desc: "This is description for first article", Content: "Hello World" },
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Homepage Endpoint")
}

func handleRequests()  {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
	
}

func main()  {
	handleRequests()
}