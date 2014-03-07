package main

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
	. "github.com/jonahgeorge/husker/controllers"
	"log"
	"net/http"
	// 	"os"
)

func main() {
	conf, _ := yaml.Open("settings.yml")
	user := to.String(conf.Get("database", "user"))
	pass := to.String(conf.Get("database", "pass"))
	name := to.String(conf.Get("database", "name"))

	db, err := sql.Open("mysql", user+":"+pass+"@/"+name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxIdleConns(100)

	// routes
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

	// auth routes
	//     http.HandleFunc("/auth/form", UserController{}.AuthForm())
	//     http.HandleFunc("/auth/email", UserController{}.AuthForm())
	//     http.HandleFunc("/auth/linkedin", UserController{}.AuthForm())

	// serve static content
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	// listen on env port
	//http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	http.ListenAndServe(":3000", nil)
}
