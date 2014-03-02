package controllers

import (
  "log"
  "net/http"
  "database/sql"
  _ "github.com/Go-SQL-Driver/MySQL"
  "github.com/hoisie/mustache"
  . "github.com/jonahgeorge/husker/models"
)

type ArticleController struct {}

func (a ArticleController) Index(db *sql.DB) string {
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

  return mustache.RenderFile("views/articles/index.mustache", data)
}

func (a ArticleController) Retrieve(db *sql.DB, req *http.Request) string {
  article, err := ArticleModel{}.RetrieveOne(db, req.FormValue("q")) 
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

  return mustache.RenderFile("views/articles/show.mustache", data)
}

func (a ArticleController) Form(db *sql.DB) string {
  return ""
}

