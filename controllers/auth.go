package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/husker/models"
	"log"
	"net/http"
)

type AuthController struct{}

func (a AuthController) SignInForm() http.HandlerFunc {
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

func (a AuthController) SignInApi(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := UserModel{}.RetrieveByEmail(db, email)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(user.Email))
		log.Println(string(user.Password))
		log.Println(password)

		err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err != nil {
			log.Println("Passwords don't match.")
			http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
		} else {
			log.Println("Passwords matched.")

			session, _ := store.Get(r, "user")
			session.Values["email"] = user.Email
			session.Values["name"] = "Jonah George"
			session.Values["id"] = user.Id

			err = session.Save(r, w)
			if err != nil {
				log.Fatal(err)
			}

			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
	}
}

func (a AuthController) SignUpForm() http.HandlerFunc {
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

func (a AuthController) SignUpApi(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Assuming password 1 and password 2 match

		email := r.FormValue("email")
		password := r.FormValue("password")

		// Assuming user doesn't exist

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 9)
		if err != nil {
			log.Fatal(err)
		}

		// Send email confirmation

		user, err := UserModel{}.Create(db, email, hashedPassword)
		if err != nil {
			log.Fatal(err)
		}

		session, _ := store.Get(r, "user")
		session.Values["email"] = user.Email
		session.Values["id"] = user.Id

		// Send to profile creation page
		http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)
	}
}
