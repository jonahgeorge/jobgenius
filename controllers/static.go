package controllers

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/sessions"
	"github.com/jonahgeorge/jobgenius.net/models"
)

type MainController struct{}

func (m MainController) Landing(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Retrieve data
		articles := models.ArticleFactory{}.GetRecent(db)
		interviews := models.InterviewFactory{}.RetrieveAll(db)
		session, err := store.Get(r, "user")

		// Catch retrieval errors and display error page
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Render template
		err = t.ExecuteTemplate(w, "landing",
			map[string]interface{}{
				"Title":      "",
				"Articles":   articles,
				"Interviews": interviews,
				"Session":    session,
			})

		// Rendering errors
		if err != nil {
			log.Println(err)
		}
	}
}

func (m MainController) About(store *sessions.CookieStore) http.HandlerFunc {
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

func (m MainController) Terms(store *sessions.CookieStore) http.HandlerFunc {
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

func (m MainController) Privacy(store *sessions.CookieStore) http.HandlerFunc {
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
