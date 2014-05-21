package controllers

import (
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/jonahgeorge/jobgenius.net/models"
)

type InterviewController struct{}

func (i InterviewController) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var interviews []models.InterviewModel
		var err error

		// Parse GET parameters
		r.ParseForm()

		session, err := store.Get(r, "user")

		if len(r.Form["title"]) > 0 {
			interviews = models.InterviewFactory{}.Filter(r.Form["title"][0])
		} else {
			interviews = models.InterviewFactory{}.RetrieveAll()
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = t.ExecuteTemplate(w, "interviews/index", map[string]interface{}{
			"Title":      "Interviews",
			"Interviews": interviews,
			"Session":    session,
		})
	}
}

func (i InterviewController) Retrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		interview := models.InterviewFactory{}.RetrieveById(params["id"])
		session, _ := store.Get(r, "user")

		err := t.ExecuteTemplate(w, "interviews/show", map[string]interface{}{
			"Title":     interview.Name,
			"Interview": interview,
			"Session":   session,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (i InterviewController) Form() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (i InterviewController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
