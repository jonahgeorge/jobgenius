package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
	"github.com/jonahgeorge/jobgenius.net/controllers"
)

func main() {
	// load settings from config file
	conf, err := yaml.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// open database connection
	db := InitDB(conf)
	defer db.Close()

	// initialize session storage
	secret := to.String(conf.Get("server", "secret"))
	store := sessions.NewCookieStore([]byte(secret))

	// init router
	router := InitRouter(db, store)

	// register router to serve all requests
	http.Handle("/", router)

	// load port from config
	port := fmt.Sprintf(":%s", to.String(conf.Get("server", "port")))

	// spin 'er up
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Initialize mux router and register necessary routes
func InitRouter(db *sql.DB, store *sessions.CookieStore) *mux.Router {
	// intialize routes muxer
	r := mux.NewRouter()

	// article routes
	r.HandleFunc("/articles", controllers.Article{}.Index(db, store))
	r.HandleFunc("/articles/{id:[0-9]+}", controllers.Article{}.Retrieve(db, store))
	//r.HandleFunc("/articles", controllers.Article{}.Form(db, store)).Methods("GET")
	r.HandleFunc("/articles", controllers.Article{}.Create(db, store)).Methods("POST")

	// interview routes
	r.HandleFunc("/interviews", controllers.Interview{}.Index(db, store))
	r.HandleFunc("/interviews/{id:[0-9]+}", controllers.Interview{}.Retrieve(db, store))
	r.HandleFunc("/interviews", controllers.Interview{}.Form(db, store))

	// account routes
	r.HandleFunc("/accounts", controllers.Account{}.Index(db, store))
	r.HandleFunc("/accounts/{id:[0-9]+}", controllers.Account{}.Retrieve(db, store))

	// user routes
	// r.HandleFunc("/signin", controllers.User{}.SignInForm(store))
	// r.HandleFunc("/signout", controllers.User{}.SignOut(store))
	// r.HandleFunc("/signup", controllers.User{}.SignUpForm(store))

	// api routes
	// r.HandleFunc("/api/signin", controllers.User{}.SignInApi(db, store))
	// r.HandleFunc("/api/signup", controllers.User{}.SignUpApi(db, store))

	// chart apis
	r.HandleFunc("/api/charts/groupwork", controllers.Chart{}.GroupWork(db, store))
	r.HandleFunc("/api/charts/fulfillment", controllers.Chart{}.Fulfillment(db, store))
	r.HandleFunc("/api/charts/breakdown", controllers.Chart{}.Breakdown(db, store))

	return r
}

// Load credentials from config and open database connection pool
func InitDB(conf *yaml.Yaml) *sql.DB {
	// load credentials from config file
	user := to.String(conf.Get("database", "username"))
	pass := to.String(conf.Get("database", "password"))
	name := to.String(conf.Get("database", "name"))

	// open database connection
	db, err := sql.Open("mysql", user+":"+pass+"@/"+name)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(100)

	return db
}
