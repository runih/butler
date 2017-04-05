package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var wiki_pages []WikiPage

type WikiPage struct {
	ID    string `json:id,omitempty`
	Title string `json:title,omitempty`
}


func GetWikiPages(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(wiki_pages)
}

func GetWikiPage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range wiki_pages {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&WikiPage{})
}

func main() {
	wiki_pages = append(wiki_pages, WikiPage{ID: "1", Title: "First Title"})
	wiki_pages = append(wiki_pages, WikiPage{ID: "2", Title: "Second Title"})
	router := mux.NewRouter()
	router.HandleFunc("/wiki_page", GetWikiPages).Methods("GET")
	router.HandleFunc("/wiki_page/{id}", GetWikiPage).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
