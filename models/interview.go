package models

import (
  "log"
  "database/sql"
  _ "github.com/Go-SQL-Driver/MySQL"
)

type InterviewTeaserModel struct {
  Id      int
  Author  string
  Date    string
  Title   string
  Content string
}

type InterviewModel struct {}
type InterviewFullModel struct {}

func (i InterviewModel) Create(db *sql.DB) error {
  return nil
}

func (i InterviewModel) RetrieveAll(db *sql.DB) ([]InterviewTeaserModel, error) {
  var teasers []InterviewTeaserModel

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
    var i InterviewTeaserModel
    err = rows.Scan(&i.Author, &i.Id, &i.Title, &i.Content, &i.Date)
    if err != nil {
      log.Fatal(err)
    }
    teasers = append(teasers, i)
  }
  return teasers, err
}

func (i InterviewModel) RetrieveOne(db *sql.DB, id string) (InterviewFullModel, error) {
  return InterviewFullModel{}, nil
}

func (i InterviewModel) Update(db *sql.DB) error {
  return nil
}

func (i InterviewModel) Delete(db *sql.DB) error {
  return nil
}
