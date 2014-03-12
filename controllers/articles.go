package controllers

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/husker/models"
	"log"
	"net/http"
)

type ArticleController struct{}

func (a ArticleController) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		articles, _ := ArticleModel{}.RetrieveAll(db)
		session, _ := store.Get(r, "user")

		data := struct {
			Title    string
			Articles []ArticleModel
			Session  *sessions.Session
		}{
			"Articles",
			articles,
			session,
		}

		if err := t.ExecuteTemplate(w, "articleIndex", data); err != nil {
			log.Fatal(err)
		}
	}
}

func (a ArticleController) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		article, _ := ArticleModel{}.RetrieveOne(db, r.FormValue("q"))
		session, _ := store.Get(r, "user")

		data := struct {
			Title   string
			Article ArticleModel
			Session *sessions.Session
		}{
			article.Title,
			article,
			session,
		}

		if err := t.ExecuteTemplate(w, "articleShow", data); err != nil {
			log.Fatal(err)
		}
	}
}

func (a ArticleController) Form(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
