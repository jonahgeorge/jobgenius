package controllers

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/jobgenius.net/models"
	"log"
	"net/http"
)

type InterviewController struct{}

func (i InterviewController) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		interviews, _ := InterviewModel{}.RetrieveAll(db)
		session, _ := store.Get(r, "user")

		data := struct {
			Title      string
			Interviews []InterviewModel
			Session    *sessions.Session
		}{
			"Interviews",
			interviews,
			session,
		}

		if err := t.ExecuteTemplate(w, "interviewIndex", data); err != nil {
			log.Fatal(err)
		}
	}
}

func (i InterviewController) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := t.ExecuteTemplate(w, "interviewShow", nil); err != nil {
			log.Fatal(err)
		}
	}
}

func (i InterviewController) Form(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
