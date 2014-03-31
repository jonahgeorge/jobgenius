package controllers

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/jobgenius.net/models"
	"net/http"
)

type Account struct{}

func (a Account) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		accounts := AccountModel{}.RetrieveAll(db)
		session, _ := store.Get(r, "user")

		err := t.ExecuteTemplate(w, "accounts/index", map[string]interface{}{
			"Title":    "Accounts",
			"Accounts": accounts,
			"Session":  session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (a Account) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		account := AccountModel{}.RetrieveById(db, params["id"])
		articles, _ := ArticleModel{}.RetrieveByAuthor(db, int(account.Id.Int64))
		interviews := InterviewModel{}.RetrieveByAuthor(db, int(account.Id.Int64))
		session, _ := store.Get(r, "user")

		err := t.ExecuteTemplate(w, "accounts/show", map[string]interface{}{
			"Title":      "Account",
			"Account":    account,
			"Articles":   articles,
			"Interviews": interviews,
			"Session":    session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
