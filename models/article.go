package models

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/russross/blackfriday"
	"html/template"
	"log"
)

type ArticleModel struct {
	Id      sql.NullInt64
	Author  sql.NullString
	Date    sql.NullString
	Title   sql.NullString
	Content template.HTML
}

func (a ArticleModel) Create(db *sql.DB, data map[string]interface{}) error {
	sql := `INSERT INTO C_ARTICLE (title, author) VALUES (?, ?)`
	_, err := db.Query(sql, data["Title"], data["Author"])
	return err
}

func (a ArticleModel) RetrieveAll(db *sql.DB) ([]ArticleModel, error) {
	var articles []ArticleModel

	sql := `SELECT U.display_name, A.aid, A.title, A.body
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
		var b []byte

		err = rows.Scan(&a.Author, &a.Id, &a.Title, &b)
		if err != nil {
			log.Fatal(err)
		}

		a.Content = template.HTML(blackfriday.MarkdownCommon(b))
		articles = append(articles, a)
	}
	return articles, err
}

func (a ArticleModel) RetrieveOne(db *sql.DB, id string) (ArticleModel, error) {

	sql := `SELECT U.display_name, A.aid, A.title, A.body, A.timestamp
               FROM C_ARTICLE AS A
                 LEFT JOIN C_USER AS U ON A.uid = U.uid
               WHERE A.published = 1 AND A.aid = ?`

	var article ArticleModel
	var b []byte

	err := db.QueryRow(sql, id).Scan(&article.Author, &article.Id, &article.Title, &b, &article.Date)
	b = blackfriday.MarkdownCommon(b)
	article.Content = template.HTML(string(b))

	return article, err
}

func (a ArticleModel) RetrieveByAuthor(db *sql.DB, id int) ([]ArticleModel, error) {
	var articles []ArticleModel

	sql := fmt.Sprintf("SELECT U.display_name, A.aid, A.title, A.body, A.timestamp FROM C_ARTICLE AS A LEFT JOIN C_USER AS U ON A.uid = U.uid WHERE A.uid = %d", id)

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var a ArticleModel
		var b []byte

		err = rows.Scan(&a.Author, &a.Id, &a.Title, &b, &a.Date)
		if err != nil {
			log.Fatal(err)
		}
		b = blackfriday.MarkdownCommon(b)
		a.Content = template.HTML(b)

		articles = append(articles, a)
	}

	return articles, err
}

func (a ArticleModel) Update(db *sql.DB) error {
	return nil
}

func (a ArticleModel) Delete(db *sql.DB) error {
	return nil
}
