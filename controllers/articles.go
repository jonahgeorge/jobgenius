package controllers

import (
  "log"
  "net/http"
  "html/template"
  "database/sql"
  . "github.com/jonahgeorge/husker/models"
  _ "github.com/Go-SQL-Driver/MySQL"
)

var t *template.Template

type ArticleController struct {}

func init() {
  t = template.Must(template.ParseGlob("templates/articles/*"))
  t = template.Must(t.ParseGlob("templates/shared/*"))
}

func (a ArticleController) Index(db *sql.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    articles, err := ArticleModel{}.RetrieveAll(db) 
    if err != nil {
      log.Fatal(err)
    }

    data := struct {
      Title    string
      Articles []ArticleModel
    }{
      "Articles",
      articles,
    }

    t.ExecuteTemplate(w, "index", data) 
  }
}

func (a ArticleController) Retrieve(db *sql.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    article, err := ArticleModel{}.RetrieveOne(db, r.FormValue("q")) 
    if err != nil {
      log.Fatal(err)
    }

    data := struct {
      Title    string
      Article   ArticleModel
    }{
      article.Title,
      article,
    }

    t.ExecuteTemplate(w, "show", data) 
  }
}

func (a ArticleController) Form(db *sql.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    t.ExecuteTemplate(w, "form", nil) 
  }
}

