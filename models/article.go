package models

import (
  "log"
  "database/sql"

  _ "github.com/Go-SQL-Driver/MySQL"
)

type Article struct {
  Id      int
  Author  string
  Date    string
  Title   string
  Content string
}

func (a Article) Create(db *sql.DB) error {
  return nil
}

func (a Article) RetrieveAll(db *sql.DB) ([]Article, error) {
  var articles []Article

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
    var a Article
    err = rows.Scan(&a.Author, &a.Id, &a.Title, &a.Content, &a.Date)
    if err != nil {
      log.Fatal(err)
    }
    articles = append(articles, a)
  }
  return articles, err
}

/*
func (a Article) Retrieve(db *sql.DB) (Article, error) {
  return nil, nil
}
*/

func (a Article) Update(db *sql.DB) error {
  return nil
}

func (a Article) Delete(db *sql.DB) error {
  return nil
}
