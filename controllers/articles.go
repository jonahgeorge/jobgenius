package controllers

import (
  "log"
  "net/http"
  "database/sql"
  // "github.com/hoisie/mustache"
  "github.com/martini-contrib/render"
  . "github.com/jonahgeorge/husker/models"
  _ "github.com/Go-SQL-Driver/MySQL"
  )

type ArticleController struct {}

func (a ArticleController) Index(ren render.Render, db *sql.DB) {
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

  ren.HTML(200, "articles/index", data)
}

func (a ArticleController) Retrieve(ren render.Render, db *sql.DB, req *http.Request) {
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

  //return mustache.RenderFile("views/articles/show.mustache", data)
  ren.HTML(200, "articles/show", data)
}

func (a ArticleController) Form(db *sql.DB) string {
  return ""
}

