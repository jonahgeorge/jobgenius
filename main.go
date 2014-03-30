package main

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
	. "github.com/jonahgeorge/jobgenius.net/controllers"
	"log"
	"net/http"
)

var db *sql.DB
var store *sessions.CookieStore
var user, pass, name, secret string

func main() {
	// load settings from config file
	conf, _ := yaml.Open("settings.yml")
	user = to.String(conf.Get("database", "user"))
	pass = to.String(conf.Get("database", "pass"))
	name = to.String(conf.Get("database", "name"))
	secret = to.String(conf.Get("session", "secret"))

	// open database connection
	db, err := sql.Open("mysql", user+":"+pass+"@/"+name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxIdleConns(100)

	// initialize session storage
	store = sessions.NewCookieStore([]byte(secret))

	// intialize gorilla router
	r := pat.New()
	r.Get("/articles", ArticleController{}.Index(db, store))
	r.Get("/article/{id:[0-9]+}", ArticleController{}.Retrieve(db, store))
	r.Get("/article", ArticleController{}.Form(db, store))
	r.Post("/article", ArticleController{}.Create(db, store))

	r.Get("/interviews", InterviewController{}.Index(db, store))
	r.Get("/interview/{id:[0-9]+}", InterviewController{}.Retrieve(db, store))
	r.Get("/interview", InterviewController{}.Form(db, store))

	r.Get("/accounts", AccountController{}.Index(db, store))
	r.Get("/account/{id:[0-9]+}", AccountController{}.Retrieve(db, store))

	r.Get("/about", StaticController{}.About(store))
	r.Get("/terms", StaticController{}.Terms(store))
	r.Get("/privacy", StaticController{}.Privacy(store))

	r.Get("/signin", UserController{}.SignInForm(store))
	r.Get("/signout", UserController{}.SignOut(store))
	r.Get("/signup", UserController{}.SignUpForm(store))

	r.Post("/api/signin", UserController{}.SignInApi(db, store))
	r.Post("/api/signup", UserController{}.SignUpApi(db, store))

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	r.Get("/", StaticController{}.Landing(db, store))

	// have our app use gorilla router
	http.Handle("/", r)

	http.ListenAndServe(":3000", nil)
}
