package controllers

import (
	"log"
	"net/http"
)

type AuthController struct{}

func (a AuthController) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Title string
		}{
			"Sign In",
		}

		err := t.ExecuteTemplate(w, "signinForm", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (a AuthController) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Title string
		}{
			"Sign Up",
		}

		err := t.ExecuteTemplate(w, "signupForm", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
