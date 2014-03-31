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

type User struct{}

func (u User) SignInForm(store *sessions.CookieStore) http.HandlerFunc {
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

func (u User) SignInApi(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		email := r.FormValue("email")
		password := r.FormValue("password")

		account := AccountModel{}.RetrieveByEmail(db, email)

		if err := bcrypt.CompareHashAndPassword([]byte(account.Password.String), []byte(password)); err != nil {
			http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
		} else {

			session, _ := store.Get(r, "user")
			session.Values["Email"] = account.Email.String
			session.Values["Name"] = account.Name.String
			session.Values["Id"] = account.Id.Int64
			session.Save(r, w)

			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
	}
}

func (u User) SignUpForm(store *sessions.CookieStore) http.HandlerFunc {
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

func (u User) SignUpApi(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Assuming password 1 and password 2 match

		email := r.FormValue("email")
		password := r.FormValue("password")

		// Assuming user doesn't exist

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 9)
		if err != nil {
			log.Fatal(err)
		}

		// Send email confirmation

		account := AccountModel{}.Create(db, email, string(hashedPass))
		var user UserModel
		user.Email = []byte(account.Email.String)
		user.Id = int(account.Id.Int64)
		user.Name = []byte("Anonymous")

		session, _ := store.Get(r, "user")
		session.Values["user"] = user
		session.Save(r, w)

		http.Redirect(w, r, "/settings", http.StatusTemporaryRedirect)
	}
}

func (u User) SignOut(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")
		session.Options.MaxAge = -1
		sessions.Save(r, w)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
