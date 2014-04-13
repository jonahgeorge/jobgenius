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

type Account struct{}

func (a Account) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		accounts := AccountModel{}.RetrieveAll(db)

		bytes, err := json.MarshalIndent(accounts, "", "\t")
		if err != nil {
			log.Printf("%s", err)
			return
		}

		fmt.Fprintf(w, "%s", bytes)
	}
}

func (a Account) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		account := AccountModel{}.RetrieveById(db, params["id"])
		articles, _ := ArticleModel{}.RetrieveByAuthor(db, int(account.Id.Int64))
		interviews := InterviewModel{}.RetrieveByAuthor(db, int(account.Id.Int64))

		bytes, err := json.MarshalIndent(map[string]interface{}{
			"account":    account,
			"articles":   articles,
			"interviews": interviews,
		}, "", "\t")

		if err != nil {
			log.Printf("%s", err)
		}

		fmt.Fprintf(w, "%s", bytes)
	}
}
