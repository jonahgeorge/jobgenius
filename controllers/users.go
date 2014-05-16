package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"code.google.com/p/go.crypto/bcrypt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	. "github.com/jonahgeorge/jobgenius.net/models"
)

type UserController struct {
}

func (u UserController) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users := UserModel{}.RetrieveAll(db)
		session, err := store.Get(r, "user")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = t.ExecuteTemplate(w, "accounts/index",
			map[string]interface{}{
				"Title":    "Accounts",
				"Accounts": users,
				"Session":  session,
			})
	}
}

func (u UserController) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		user := UserModel{}.RetrieveById(db, params["id"])
		articles, err := ArticleFactory{}.RetrieveByAuthor(db, *user.Id)
		interviews := InterviewFactory{}.RetrieveByAuthor(db, *user.Id)
		session, err := store.Get(r, "user")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = t.ExecuteTemplate(w, "accounts/show",
			map[string]interface{}{
				"Title":      "Account",
				"Account":    user,
				"Articles":   articles,
				"Interviews": interviews,
				"Session":    session,
			})
	}
}

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

		user := UserModel{}.RetrieveByEmail(db, email)

		err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
		if err != nil {
			http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
			return
		}

		session, _ := store.Get(r, "user")
		session.Values["Email"] = user.Email
		session.Values["DisplayName"] = user.DisplayName
		session.Values["Id"] = user.Id
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
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

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 9)
		if err != nil {
			log.Fatal(err)
		}

		// Send email confirmation

		account := UserModel{}.Create(db, email, string(hashedPass))

		session, _ := store.Get(r, "user")
		session.Values["user"] = account
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
