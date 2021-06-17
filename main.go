package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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
		Article{Title: "First Article Title", Desc: "This is description for first article", Content: "Hello World" },
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Homepage Endpoint")
}

func postArticles(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "POST endpoint worked")
}

func handleRequests()  {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	//can restrict endpoint to only be accessible with GET method
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	//same enpoint as above but using different method
	myRouter.HandleFunc("/articles", postArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
	
}

func main()  {
	handleRequests()
}