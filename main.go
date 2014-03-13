package main

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
	. "github.com/jonahgeorge/husker/controllers"
	"log"
	"net/http"
)

var db *sql.DB
var store *sessions.CookieStore
var user, pass, name, secret string

func main() {
	conf, _ := yaml.Open("settings.yml")
	user = to.String(conf.Get("database", "user"))
	pass = to.String(conf.Get("database", "pass"))
	name = to.String(conf.Get("database", "name"))
	secret = to.String(conf.Get("session", "secret"))

	db, err := sql.Open("mysql", user+":"+pass+"@/"+name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxIdleConns(100)

	store = sessions.NewCookieStore([]byte(secret))

	r := mux.NewRouter()
	r.HandleFunc("/articles", ArticleController{}.Index(db, store))
	r.HandleFunc("/article/{id:[0-9]+}", ArticleController{}.Retrieve(db, store))
	r.HandleFunc("/interviews", InterviewController{}.Index(db, store))
	r.HandleFunc("/interview/{id:[0-9]+}", InterviewController{}.Retrieve(db, store))
	r.HandleFunc("/accounts", AccountController{}.Index(db, store))
	r.HandleFunc("/account/{id:[0-9]+}", AccountController{}.Retrieve(db, store))
	r.HandleFunc("/", StaticController{}.Landing(db, store))
	r.HandleFunc("/about", StaticController{}.About(store))
	r.HandleFunc("/terms", StaticController{}.Terms(store))
	r.HandleFunc("/privacy", StaticController{}.Privacy(store))
	r.HandleFunc("/signin", UserController{}.SignInForm(store))
	r.HandleFunc("/signout", UserController{}.SignOut(store))
	r.HandleFunc("/signup", UserController{}.SignUpForm(store))
	r.HandleFunc("/api/signin", UserController{}.SignInApi(db, store))
	r.HandleFunc("/api/signup", UserController{}.SignUpApi(db, store))

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.Handle("/", r)

	// listen on env port
	http.ListenAndServe(":3000", nil)
}
