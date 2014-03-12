package controllers

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/husker/models"
	"log"
	"net/http"
)

type StaticController struct{}

func (s StaticController) Landing(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		accounts, _ := AccountModel{}.RetrieveAll(db)
		articles, _ := ArticleModel{}.RetrieveAll(db)
		interviews, _ := InterviewModel{}.RetrieveAll(db)
		session, _ := store.Get(r, "user")

		data := struct {
			Title      string
			Accounts   []AccountModel
			Articles   []ArticleModel
			Interviews []InterviewModel
			Session    *sessions.Session
		}{
			"Welcome",
			accounts,
			articles,
			interviews,
			session,
		}

		if err := t.ExecuteTemplate(w, "landingTemplate", data); err != nil {
			log.Fatal(err)
		}
	}
}

func (s StaticController) About(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		data := struct {
			Title   string
			Session *sessions.Session
		}{
			"Terms and Conditions",
			session,
		}

		if err := t.ExecuteTemplate(w, "aboutTemplate", data); err != nil {
			log.Fatal(err)
		}
	}
}

func (s StaticController) Terms(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		data := struct {
			Title   string
			Session *sessions.Session
		}{
			"Terms and Conditions",
			session,
		}

		if err := t.ExecuteTemplate(w, "termsTemplate", data); err != nil {
			log.Fatal(err)
		}
	}
}

func (s StaticController) Privacy(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		data := struct {
			Title   string
			Session *sessions.Session
		}{
			"Terms and Conditions",
			session,
		}

		if err := t.ExecuteTemplate(w, "privacyTemplate", data); err != nil {
			log.Fatal(err)
		}
	}
}
