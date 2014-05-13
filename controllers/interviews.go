package controllers

import (
	"database/sql"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jonahgeorge/jobgenius.net/models"
)

type InterviewController struct{}

func (i InterviewController) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		interviews := models.InterviewFactory{}.RetrieveAll(db)
		session, _ := store.Get(r, "user")

		err := t.ExecuteTemplate(w, "interviews/index", map[string]interface{}{
			"Title":      "Interviews",
			"Interviews": interviews,
			"Session":    session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (i InterviewController) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		interview := models.InterviewFactory{}.RetrieveById(db, params["id"])
		session, _ := store.Get(r, "user")

		err := t.ExecuteTemplate(w, "interviews/show", map[string]interface{}{
			"Title":     interview.Name.String,
			"Interview": interview,
			"Session":   session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (i InterviewController) Form(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (i InterviewController) Create(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
