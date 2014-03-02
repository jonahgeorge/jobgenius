package controllers

import (
  "log"
  "net/http"
  "database/sql"

  _ "github.com/Go-SQL-Driver/MySQL"
  "github.com/hoisie/mustache"
//  "github.com/codegangsta/martini"
  "github.com/jonahgeorge/husker/models"
)

type ArticleController struct {}

func (a ArticleController) Index(db *sql.DB) string {
  articles, err := models.Article{}.RetrieveAll(db) 
  if err != nil {
    log.Fatal(err)
  }
  data := struct {
    Articles []models.Article
  }{
    articles,
  }
  return mustache.RenderFile("views/articles/index.mustache", data)
}

func (a ArticleController) Retrieve(db *sql.DB, req *http.Request) string {
  return ""
}

func (a ArticleController) Form(db *sql.DB) string {
  return ""
}

