package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/jobgenius.net/models"
	"log"
	"net/http"
)

type UserController struct{}

func (u UserController) SignInForm(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		data := struct {
			Title   string
			Session *sessions.Session
		}{
			"Sign In",
			session,
		}

		if err := t.ExecuteTemplate(w, "signinForm", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (u UserController) SignInApi(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		account := AccountModel{}.RetrieveByEmail(db, email)

		if err := bcrypt.CompareHashAndPassword(account.Password, []byte(password)); err != nil {
			http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
		} else {

			var user UserModel
			user.Email = account.Email
			user.Name = account.Name
			user.Id = account.Id

			session, _ := store.Get(r, "user")
			session.Values["user"] = user
			session.Save(r, w)

			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
	}
}

func (u UserController) SignUpForm(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		data := struct {
			Title   string
			Session *sessions.Session
		}{
			"Sign Up",
			session,
		}

		if err := t.ExecuteTemplate(w, "signupForm", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (u UserController) SignUpApi(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
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

		account := AccountModel{}.Create(db, email, hashedPassword)
		var user UserModel
		user.Email = account.Email
		user.Id = account.Id
		user.Name = []byte("Anonymous")

		session, _ := store.Get(r, "user")
		session.Values["user"] = user
		session.Save(r, w)

		http.Redirect(w, r, "/settings", http.StatusTemporaryRedirect)
	}
}

func (u UserController) SignOut(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")
		session.Options.MaxAge = -1
		sessions.Save(r, w)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
