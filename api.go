package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getWikiPages(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(wikiPages)
}

func getWikiPage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	page := fetchWikiPage(params["id"])
	if page != nil {
		json.NewEncoder(w).Encode(&page)
		return
	}
	json.NewEncoder(w).Encode(&WikiPage{})
}
