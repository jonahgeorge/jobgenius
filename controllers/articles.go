package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/jobgenius.net/models"
	"github.com/russross/blackfriday"
	"html/template"
	"log"
	"net/http"
)

type Article struct{}

// Handles the rendering of all articles to the index page
func (a Article) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
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
func (a Article) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		article, err := ArticleModel{}.RetrieveById(db, params["id"])
		session, err := store.Get(r, "user")
		mrkdwn := template.HTML(blackfriday.MarkdownCommon([]byte(article.Content.String)))

		err = t.ExecuteTemplate(w, "articleShow", map[string]interface{}{
			"Title":   article.Title.String,
			"Date":    article.Date.String,
			"Content": mrkdwn,
			"Session": session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Handles the rendering of the new article form page
func (a Article) Form(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "user")
		log.Printf("%+v", session)
		log.Printf("%s", session.Values["Id"])

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
func (a Article) Create(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		id, err := ArticleModel{}.Create(db, map[string]interface{}{
			"AuthorId": session.Values["Id"],
			"Title":    r.FormValue("title"),
			"Content":  r.FormValue("content"),
		})

		log.Printf("%s", id)

		if err != nil {
			log.Printf("%s", err)
			http.Redirect(w, r, "/articles", http.StatusTemporaryRedirect)
		} else {
			url := fmt.Sprintf("/article/%d", id)
			log.Printf("%s", url)
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		}
	}
}
