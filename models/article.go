package models

import (
// "fmt"
  "log"
  "database/sql"
  _ "github.com/Go-SQL-Driver/MySQL"
)

type ArticleModel struct {
  Id      int
  Author  string
  Date    string
  Title   string
  Content string
}

func (a ArticleModel) Create(db *sql.DB) error {
  return nil
}

func (a ArticleModel) RetrieveAll(db *sql.DB) ([]ArticleModel, error) {
  var articles []ArticleModel

  sql := `SELECT U.display_name, A.aid, A.title, A.body, A.timestamp
          FROM C_ARTICLE AS A
            LEFT JOIN C_USER AS U ON A.uid = U.uid
          WHERE A.published = 1`

  rows, err := db.Query(sql)
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
    var a ArticleModel
    err = rows.Scan(&a.Author, &a.Id, &a.Title, &a.Content, &a.Date)
    if err != nil {
      log.Fatal(err)
    }
    articles = append(articles, a)
  }
  return articles, err
}

func (a ArticleModel) RetrieveOne(db *sql.DB, id string) (ArticleModel, error) {

  sql := `SELECT U.display_name, A.aid, A.title, A.body, A.timestamp
          FROM C_ARTICLE AS A
            LEFT JOIN C_USER AS U ON A.uid = U.uid
          WHERE A.published = 1 AND A.aid = ` + id

  var article ArticleModel
  err := db.QueryRow(sql).Scan(&article.Author, &article.Id, &article.Title, &article.Content, &article.Date)
  if err != nil {
    log.Fatal(err)
  }

  return article, err
}

func (a ArticleModel) Update(db *sql.DB) error {
  return nil
}

func (a ArticleModel) Delete(db *sql.DB) error {
  return nil
}
