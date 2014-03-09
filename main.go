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

func main() {
	conf, _ := yaml.Open("settings.yml")
	user := to.String(conf.Get("database", "user"))
	pass := to.String(conf.Get("database", "pass"))
	name := to.String(conf.Get("database", "name"))
	secret := to.String(conf.Get("session", "secret"))

	db, err := sql.Open("mysql", user+":"+pass+"@/"+name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxIdleConns(100)

	store := sessions.NewCookieStore([]byte(secret))

	http.HandleFunc("/articles", ArticleController{}.Index(db))
	http.HandleFunc("/article", ArticleController{}.Retrieve(db))

	http.HandleFunc("/interviews", InterviewController{}.Index(db))
	http.HandleFunc("/interview", InterviewController{}.Retrieve(db))

	http.HandleFunc("/users", UserController{}.Index(db))
	http.HandleFunc("/user", UserController{}.Retrieve(db))

	http.HandleFunc("/", StaticController{}.Landing(db))
	http.HandleFunc("/about", StaticController{}.About())
	http.HandleFunc("/terms", StaticController{}.Terms())
	http.HandleFunc("/privacy", StaticController{}.Privacy())

	http.HandleFunc("/signin", AuthController{}.SignInForm())
	http.HandleFunc("/api/signin", AuthController{}.SignInApi(db, store))
	http.HandleFunc("/signup", AuthController{}.SignUpForm())
	http.HandleFunc("/api/signup", AuthController{}.SignUpApi(db, store))
	//     http.HandleFunc("/auth/email", UserController{}.AuthForm())
	//     http.HandleFunc("/auth/linkedin", UserController{}.AuthForm())

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
