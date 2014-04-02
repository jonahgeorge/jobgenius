package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/sessions"
	"github.com/jonahgeorge/jobgenius.net/models/blocks"
	"net/http"
)

type Chart struct{}

func (c Chart) GroupWork(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		data := blocks.GroupworkChart{}.RetrieveById(db, params.Get("i"))

		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprintf(w, string(b))
	}
}

func (c Chart) Fulfillment(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		data := blocks.FulfillmentChart{}.RetrieveById(db, params.Get("i"))

		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprintf(w, string(b))
	}
}

func (c Chart) Breakdown(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		data := blocks.BreakdownChart{}.RetrieveById(db, params.Get("i"))

		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprintf(w, string(b))
	}
}
