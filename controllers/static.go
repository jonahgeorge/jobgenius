package controllers

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	. "github.com/jonahgeorge/husker/models"
	"log"
	"net/http"
)

type StaticController struct{}

func (s StaticController) Landing(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := UserModel{}.RetrieveAll(db)
		articles, err := ArticleModel{}.RetrieveAll(db)
		interviews, err := InterviewModel{}.RetrieveAll(db)
		if err != nil {
			log.Fatal(err)
		}
		signedIn := false

		data := struct {
			Title      string
			Users      []UserModel
			Articles   []ArticleModel
			Interviews []InterviewModel
		}{
			"Welcome",
			users,
			articles,
			interviews,
		}

		if signedIn == false {
			err := t.ExecuteTemplate(w, "landingTemplate", data)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (s StaticController) About() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "aboutTemplate", nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s StaticController) Terms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "termsTemplate", nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s StaticController) Privacy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "privacyTemplate", nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
