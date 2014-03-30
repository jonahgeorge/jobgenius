package controllers

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/jobgenius.net/models"
	"net/http"
)

type ArticleController struct{}

// Handles the rendering of all articles to the index page
func (a ArticleController) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		articles, err := ArticleModel{}.RetrieveAll(db)
		session, err := store.Get(r, "user")

		err = t.ExecuteTemplate(w, "articleIndex", map[string]interface{}{
			"Title":    "Articles",
			"Articles": articles,
			"Session":  session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Handles the retrieval and rendering of a single article by the 'id' parameter
func (a ArticleController) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		article, err := ArticleModel{}.RetrieveOne(db, params["id"])
		session, err := store.Get(r, "user")

		err = t.ExecuteTemplate(w, "articleShow", map[string]interface{}{
			"Title":   article.Title.String,
			"Article": article,
			"Session": session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Handles the rendering of the new article form page
func (a ArticleController) Form(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "user")

		err = t.ExecuteTemplate(w, "articleForm", map[string]interface{}{
			"Title":   "New Article",
			"Session": session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Handles the creation of articles from the article form page
func (a ArticleController) Create(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		ArticleModel{}.Create(db, map[string]interface{}{
			"Author": session.Values["Name"],
			"Title":  r.FormValue("title"),
		})
	}
}
