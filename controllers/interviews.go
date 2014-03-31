package controllers

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/jobgenius.net/models"
	"net/http"
)

type Interview struct{}

func (i Interview) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		interviews := InterviewModel{}.RetrieveAll(db)
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

func (i Interview) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		interview := InterviewModel{}.RetrieveById(db, params["id"])
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

func (i Interview) Form(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
