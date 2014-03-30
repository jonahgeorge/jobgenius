package controllers

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/jobgenius.net/models"
	"log"
	"net/http"
)

type AccountController struct{}

func (a AccountController) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accounts, _ := AccountModel{}.RetrieveAll(db)

		session, _ := store.Get(r, "user")

		data := struct {
			Title    string
			Accounts []AccountModel
			Session  *sessions.Session
		}{
			"Accounts",
			accounts,
			session,
		}

		if err := t.ExecuteTemplate(w, "accountIndex", data); err != nil {
			log.Fatal(err)
		}
	}
}

func (a AccountController) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		account, _ := AccountModel{}.RetrieveById(db, params["id"])
		articles, _ := ArticleModel{}.RetrieveByAuthor(db, account.Id)
		interviews, _ := InterviewModel{}.RetrieveByAuthor(db, account.Id)

		session, _ := store.Get(r, "user")

		data := struct {
			Title      string
			Account    AccountModel
			Articles   []ArticleModel
			Interviews []InterviewModel
			Session    *sessions.Session
		}{
			"Account",
			account,
			articles,
			interviews,
			session,
		}

		if err := t.ExecuteTemplate(w, "accountShow", data); err != nil {
			log.Fatal(err)
		}
	}
}
