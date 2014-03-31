package main

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
	"github.com/jonahgeorge/jobgenius.net/controllers"
	"log"
	"net/http"
)

func main() {
	// load settings from config file
	conf, _ := yaml.Open("settings.yml")
	user := to.String(conf.Get("database", "user"))
	pass := to.String(conf.Get("database", "pass"))
	name := to.String(conf.Get("database", "name"))
	secret := to.String(conf.Get("session", "secret"))
	port := to.String(conf.Get("port"))

	// open database connection
	db, err := sql.Open("mysql", user+":"+pass+"@/"+name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxIdleConns(100)

	// initialize session storage
	store := sessions.NewCookieStore([]byte(secret))

	// intialize routes muxer
	r := mux.NewRouter()

	// article routes
	r.HandleFunc("/articles", controllers.Article{}.Index(db, store))
	r.HandleFunc("/article/{id:[0-9]+}", controllers.Article{}.Retrieve(db, store))
	r.HandleFunc("/article", controllers.Article{}.Form(db, store)).Methods("GET")
	r.HandleFunc("/article", controllers.Article{}.Create(db, store)).Methods("POST")

	// interview routes
	r.HandleFunc("/interviews", controllers.Interview{}.Index(db, store))
	r.HandleFunc("/interview/{id:[0-9]+}", controllers.Interview{}.Retrieve(db, store))
	r.HandleFunc("/interview", controllers.Interview{}.Form(db, store))

	// account routes
	r.HandleFunc("/accounts", controllers.Account{}.Index(db, store))
	r.HandleFunc("/account/{id:[0-9]+}", controllers.Account{}.Retrieve(db, store))

	// static page routes
	r.HandleFunc("/about", controllers.Static{}.About(store))
	r.HandleFunc("/terms", controllers.Static{}.Terms(store))
	r.HandleFunc("/privacy", controllers.Static{}.Privacy(store))

	// user routes
	r.HandleFunc("/signin", controllers.User{}.SignInForm(store))
	r.HandleFunc("/signout", controllers.User{}.SignOut(store))
	r.HandleFunc("/signup", controllers.User{}.SignUpForm(store))

	// api routes
	r.HandleFunc("/api/signin", controllers.User{}.SignInApi(db, store))
	r.HandleFunc("/api/signup", controllers.User{}.SignUpApi(db, store))

	// static resource files
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// route path
	r.HandleFunc("/", controllers.Static{}.Landing(db, store))

	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
