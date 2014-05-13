package controllers

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/sessions"
	"github.com/jonahgeorge/jobgenius.net/models"
)

type Static struct{}

func (s Static) Landing(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		accounts := models.AccountModel{}.RetrieveAll(db)
		articles, _ := models.ArticleModel{}.RetrieveAll(db)
		interviews := models.InterviewFactory{}.RetrieveAll(db)
		session, _ := store.Get(r, "user")

		err := t.ExecuteTemplate(w, "landing", map[string]interface{}{
			"Title":      "",
			"Accounts":   accounts,
			"Articles":   articles,
			"Interviews": interviews,
			"Session":    session,
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s Static) About(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		err := t.ExecuteTemplate(w, "aboutTemplate", map[string]interface{}{
			"Title":   "About",
			"Session": session,
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s Static) Terms(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		err := t.ExecuteTemplate(w, "termsTemplate", map[string]interface{}{
			"Title":   "Terms and Conditions",
			"Session": session,
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s Static) Privacy(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		err := t.ExecuteTemplate(w, "privacyTemplate", map[string]interface{}{
			"Title":   "Privacy Policy",
			"Session": session,
		})

		if err != nil {
			log.Fatal(err)
		}

	}
}
