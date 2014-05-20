package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	. "github.com/jonahgeorge/jobgenius.net/controllers"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
)

func initRouter() {

}

func main() {
	// load settings from config file
	conf, _ := yaml.Open("config.yml")
	user := to.String(conf.Get("database", "username"))
	pass := to.String(conf.Get("database", "password"))
	name := to.String(conf.Get("database", "name"))
	secret := to.String(conf.Get("server", "secret"))
	port := to.String(conf.Get("server", "port"))

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
	r.HandleFunc("/articles", ArticleController{}.Index(db, store)).Methods("GET")
	r.HandleFunc("/articles/{id:[0-9]+}", ArticleController{}.Retrieve(db, store)).Methods("GET")
	// r.HandleFunc("/articles/{id:[0-9]+}", ArticleController{}.Retrieve(db, store)).Methods("POST")
	r.HandleFunc("/articles/new", ArticleController{}.Form(db, store)).Methods("GET")
	r.HandleFunc("/articles", ArticleController{}.Create(db, store)).Methods("POST")

	// interview routes
	r.HandleFunc("/interviews", InterviewController{}.Index(db, store)).Methods("GET")
	r.HandleFunc("/interviews/{id:[0-9]+}", InterviewController{}.Retrieve(db, store)).Methods("GET")
	r.HandleFunc("/interviews/new", InterviewController{}.Form(db, store)).Methods("GET")
	r.HandleFunc("/interviews/new", InterviewController{}.Create(db, store)).Methods("POST")

	// account routes
	r.HandleFunc("/accounts", UserController{}.Index(db, store))
	r.HandleFunc("/accounts/{id:[0-9]+}", UserController{}.Retrieve(db, store))

	// static page routes
	r.HandleFunc("/about", MainController{}.About(store))
	r.HandleFunc("/terms", MainController{}.Terms(store))
	r.HandleFunc("/privacy", MainController{}.Privacy(store))
	r.HandleFunc("/", MainController{}.Landing(db, store))

	// user routes
	r.HandleFunc("/signin", UserController{}.SignInForm(store)).Methods("GET")
	r.HandleFunc("/signin", UserController{}.SignInApi(db, store)).Methods("POST")
	r.HandleFunc("/signout", UserController{}.SignOut(store)).Methods("GET")
	r.HandleFunc("/signup", UserController{}.SignUpForm(store)).Methods("GET")
	r.HandleFunc("/signup", UserController{}.SignUpApi(db, store)).Methods("POST")

	// api routes
	r.HandleFunc("/api/charts/groupwork", Chart{}.GroupWork(db, store))
	r.HandleFunc("/api/charts/fulfillment", Chart{}.Fulfillment(db, store))
	r.HandleFunc("/api/charts/breakdown", Chart{}.Breakdown(db, store))

	// static resource files
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("vendor"))))

	// register gorrilla router as root
	http.Handle("/", r)

	if err = http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
