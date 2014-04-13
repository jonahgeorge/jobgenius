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

type Interview struct{}

func (i Interview) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		interviews := InterviewModel{}.RetrieveAll(db)

		bytes, err := json.MarshalIndent(interviews, "", "\t")
		if err != nil {
			log.Printf("%s", err)
			return
		}

		fmt.Fprintf(w, "%s", bytes)
	}
}

func (i Interview) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		interview := InterviewModel{}.RetrieveById(db, params["id"])

		bytes, err := json.MarshalIndent(interview, "", "\t")
		if err != nil {
			log.Printf("%s", err)
			return
		}

		fmt.Fprintf(w, "%s", bytes)
	}
}

func (i Interview) Form(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
