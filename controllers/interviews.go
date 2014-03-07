package controllers

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	. "github.com/jonahgeorge/husker/models"
	"log"
	"net/http"
)

type InterviewController struct{}

func (i InterviewController) Index(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		interviews, err := InterviewModel{}.RetrieveAll(db)
		if err != nil {
			log.Fatal(err)
		}

		data := struct {
			Title      string
			Interviews []InterviewModel
		}{
			"Interviews",
			interviews,
		}

		err = t.ExecuteTemplate(w, "interviewIndex", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (i InterviewController) Retrieve(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "interviewShow", nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (i InterviewController) Form(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
