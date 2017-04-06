package main

import (
	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/wiki_page", getWikiPages).Methods("GET")
	r.HandleFunc("/wiki_page/{id}", getWikiPage).Methods("GET")
	return r
}
