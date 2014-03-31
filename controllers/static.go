package controllers

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/jobgenius.net/models"
	"log"
	"net/http"
)

type Static struct{}

func (s Static) Landing(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		accounts := AccountModel{}.RetrieveAll(db)
		articles, _ := ArticleModel{}.RetrieveAll(db)
		interviews := InterviewModel{}.RetrieveAll(db)
		session, _ := store.Get(r, "user")

		err := t.ExecuteTemplate(w, "landingTemplate", map[string]interface{}{
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
