package controllers

import (
	"log"
	"net/http"
)

type StaticController struct{}

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
