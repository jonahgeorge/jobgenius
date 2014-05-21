package controllers

import (
	"html/template"
	"log"

	"github.com/gorilla/sessions"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
)

var (
	t     *template.Template
	store *sessions.CookieStore
)

func init() {
	// parse templates
	t = template.Must(t.ParseGlob("views/_templates/*.html"))
	t = template.Must(t.ParseGlob("views/articles/*.html"))
	t = template.Must(t.ParseGlob("views/interviews/*.html"))
	t = template.Must(t.ParseGlob("views/interviews/components/*.html"))
	t = template.Must(t.ParseGlob("views/accounts/*.html"))
	t = template.Must(t.ParseGlob("views/users/*.html"))
	t = template.Must(t.ParseGlob("views/*.html"))

	// Load config file
	conf, err := yaml.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// grab server secret from config file
	secret := to.String(conf.Get("server", "secret"))

	// initialize session storage
	store = sessions.NewCookieStore([]byte(secret))
}
