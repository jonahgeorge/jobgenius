package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/context"
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

	http.HandleFunc("/articles", ArticleController{}.Index(db, store))
	http.HandleFunc("/article", ArticleController{}.Retrieve(db, store))
	http.HandleFunc("/interviews", InterviewController{}.Index(db, store))
	http.HandleFunc("/interview", InterviewController{}.Retrieve(db, store))
	http.HandleFunc("/accounts", AccountController{}.Index(db, store))
	http.HandleFunc("/account", AccountController{}.Retrieve(db, store))
	http.HandleFunc("/", StaticController{}.Landing(db, store))
	http.HandleFunc("/about", StaticController{}.About(store))
	http.HandleFunc("/terms", StaticController{}.Terms(store))
	http.HandleFunc("/privacy", StaticController{}.Privacy(store))
	http.HandleFunc("/signin", UserController{}.SignInForm(store))
	http.HandleFunc("/signout", UserController{}.SignOut(store))
	http.HandleFunc("/signup", UserController{}.SignUpForm(store))
	http.HandleFunc("/api/signin", UserController{}.SignInApi(db, store))
	http.HandleFunc("/api/signup", UserController{}.SignUpApi(db, store))

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "user")
		//email := fmt.Sprintf("%s", session.Values["email"])
		for _, value := range session.Values {
			fmt.Fprintf(w, "%s", value)
		}
	})

	// listen on env port
	http.ListenAndServe(":3000", context.ClearHandler(http.DefaultServeMux))
}
