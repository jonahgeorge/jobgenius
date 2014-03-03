package main

import (
  "os"
  "log"
  "net/http"
  "database/sql"
  "github.com/gosexy/yaml"
  "github.com/gosexy/to"
  _ "github.com/Go-SQL-Driver/MySQL"
  . "github.com/jonahgeorge/husker/controllers"
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

  // routes
  http.HandleFunc("/articles", ArticleController{}.Index(db))
  http.HandleFunc("/article", ArticleController{}.Retrieve(db))

  // listen on environmetn port
  http.ListenAndServe(":" + os.Getenv("PORT"), nil)
  // http.ListenAndServe(":3000", nil)
}
