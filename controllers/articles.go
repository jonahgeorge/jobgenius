package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/jobgenius.net/models"
)

type Article struct{}

// Handles the rendering of all articles to the index page
func (a Article) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		articles, err := ArticleModel{}.RetrieveAll(db)
		if err != nil {
			log.Printf("%s", err)
			return
		}

		bytes, err := json.MarshalIndent(articles, "", "\t")
		if err != nil {
			log.Printf("%s", err)
			return
		}

		fmt.Fprintf(w, "%s", bytes)
	}
}

// Handles the retrieval and rendering of a single article by the 'id' parameter
func (a Article) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		article, err := ArticleModel{}.RetrieveById(db, params["id"])
		if err != nil {
			log.Printf("%s", err)
			return
		}

		bytes, err := json.MarshalIndent(article, "", "\t")
		if err != nil {
			log.Printf("%s", err)
			return
		}

		fmt.Fprintf(w, "%s", bytes)
	}
}

// Handles the rendering of the new article form page
func (a Article) Form(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Handles the creation of articles from the article form page
func (a Article) Create(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		_, err := ArticleModel{}.Create(db, map[string]interface{}{
			"AuthorId": session.Values["Id"],
			"Title":    r.FormValue("title"),
			"Content":  r.FormValue("content"),
		})

		if err != nil {
			log.Printf("%s", err)
			fmt.Fprintf(w, "%s", "Failure")
		} else {
			log.Printf("%s", err)
			fmt.Fprintf(w, "%s", "Success")
		}
	}
}
