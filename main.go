package main

import (
	"log"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	. "github.com/jonahgeorge/jobgenius.net/controllers"

	"github.com/gorilla/mux"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
)

func main() {
	// intialize routes muxer
	r := mux.NewRouter()

	// article routes
	r.HandleFunc("/articles", ArticleController{}.Index()).Methods("GET")
	r.HandleFunc("/articles", ArticleController{}.Create()).Methods("POST")
	r.HandleFunc("/articles/new", ArticleController{}.Form()).Methods("GET")
	r.HandleFunc("/articles/{id:[0-9]+}", ArticleController{}.Retrieve()).Methods("GET")
	r.HandleFunc("/articles/{id:[0-9]+}/publish", ArticleController{}.Publish()).Methods("POST")
	r.HandleFunc("/articles/{id:[0-9]+}/edit", ArticleController{}.Edit()).Methods("GET")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", ArticleController{}.Delete()).Methods("POST")

	// interview routes
	r.HandleFunc("/interviews", InterviewController{}.Index()).Methods("GET")
	r.HandleFunc("/interviews/{id:[0-9]+}", InterviewController{}.Retrieve()).Methods("GET")
	r.HandleFunc("/interviews/new", InterviewController{}.Form()).Methods("GET")
	r.HandleFunc("/interviews/new", InterviewController{}.Create()).Methods("POST")

	// account routes
	r.HandleFunc("/accounts", UserController{}.Index())
	r.HandleFunc("/accounts/{id:[0-9]+}", UserController{}.Retrieve())

	// static page routes
	r.HandleFunc("/about", MainController{}.About())
	r.HandleFunc("/terms", MainController{}.Terms())
	r.HandleFunc("/privacy", MainController{}.Privacy())
	r.HandleFunc("/", MainController{}.Landing())

	// user routes
	r.HandleFunc("/signin", UserController{}.SignInForm()).Methods("GET")
	r.HandleFunc("/signin", UserController{}.SignInApi()).Methods("POST")
	r.HandleFunc("/signout", UserController{}.SignOut()).Methods("GET")
	r.HandleFunc("/signup", UserController{}.SignUpForm()).Methods("GET")
	r.HandleFunc("/signup", UserController{}.SignUpApi()).Methods("POST")

	// api routes
	r.HandleFunc("/api/charts/groupwork", Chart{}.GroupWork())
	r.HandleFunc("/api/charts/fulfillment", Chart{}.Fulfillment())
	r.HandleFunc("/api/charts/breakdown", Chart{}.Breakdown())

	// static resource files
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("vendor"))))

	// register gorrilla router as root
	http.Handle("/", r)

	// Load config file
	conf, err := yaml.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	port := to.String(conf.Get("server", "port"))
	if err = http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
