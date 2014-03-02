package main

import (
  "log"
  "database/sql"

  _ "github.com/Go-SQL-Driver/MySQL"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/gzip"
  "github.com/gosexy/yaml"
  "github.com/gosexy/to"

  "github.com/jonahgeorge/husker/controllers"
)

func main() {

  conf, _ := yaml.Open("settings.yaml")
  user := to.String(conf.Get("database", "user"))
  pass := to.String(conf.Get("database", "pass"))
  name := to.String(conf.Get("database", "name"))

  db, err := sql.Open("mysql", user + ":" + pass + "@/" + name)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()
  db.SetMaxIdleConns(100)

  m := martini.Classic()

  // gzip all responses
  m.Use(gzip.All())
  // inject db into all handlers
  m.Map(db)

  // routes
  m.Get("/articles", controllers.ArticleController{}.Index)

  m.Run()
}