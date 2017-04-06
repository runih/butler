package main

import (
	"database/sql"
	"log"
	"net/http"
)

var wikiPages []WikiPage

var db *sql.DB

func main() {
	wikiPages = append(wikiPages, WikiPage{ID: "1", Title: "First Title"})
	wikiPages = append(wikiPages, WikiPage{ID: "2", Title: "Second Title"})
	db = openDB()

	log.Fatal(http.ListenAndServe(":12345", router()))
	db.Close()
}
