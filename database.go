package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func openDB() *sql.DB {
	db, err := sql.Open("postgres", "user=additude dbname=wikipedia sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func fetchWikiPage(id string) *WikiPage {
	wikiPage := WikiPage{ID: id}
	err := db.QueryRow("SELECT title, text, created_at FROM pages WHERE id = $1;", wikiPage.ID).Scan(&wikiPage.Title, &wikiPage.Text, &wikiPage.CreatedAt)
	if err != nil {
		return &WikiPage{}
	}
	return &wikiPage
}
