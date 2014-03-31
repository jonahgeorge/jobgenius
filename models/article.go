package models

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type ArticleModel struct {
	Id      sql.NullInt64
	Author  sql.NullString
	Date    sql.NullString
	Title   sql.NullString
	Content sql.NullString
}

// Create an article
func (a ArticleModel) Create(db *sql.DB, data map[string]interface{}) (int64, error) {
	sql := `INSERT INTO C_ARTICLE (title, uid, body, published) VALUES (?, ?, ?, 1)`
	result, err := db.Exec(sql, data["Title"], data["AuthorId"], data["Content"])
	if err != nil {
		log.Printf("%s", err)
	}

	id, err := result.LastInsertId()

	return id, err
}

// Retrieve all articles
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
		var article ArticleModel

		err = rows.Scan(&article.Author, &article.Id, &article.Title, &article.Content)
		if err != nil {
			log.Fatal(err)
		}

		articles = append(articles, article)
	}
	return articles, err
}

// Retrieve one article by article id (primary key)
func (a ArticleModel) RetrieveById(db *sql.DB, id string) (ArticleModel, error) {

	sql := `SELECT 
              U.display_name, A.aid, A.title, A.body, A.timestamp
            FROM 
              C_ARTICLE AS A
            LEFT JOIN 
              C_USER AS U ON A.uid = U.uid
            WHERE 
              A.published = 1 AND A.aid = ?`

	var article ArticleModel
	row := db.QueryRow(sql, id)
	err := row.Scan(&article.Author, &article.Id, &article.Title, &article.Content, &article.Date)
	return article, err
}

// Retrieve a slice of articles authored by the user id parameter
func (a ArticleModel) RetrieveByAuthor(db *sql.DB, id int) ([]ArticleModel, error) {
	var articles []ArticleModel

	sql := `SELECT 
              U.display_name, A.aid, A.title, A.body, A.timestamp 
            FROM 
              C_ARTICLE AS A 
            LEFT JOIN 
              C_USER AS U ON A.uid = U.uid 
            WHERE 
              A.uid = ?`

	rows, err := db.Query(sql, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var article ArticleModel
		err = rows.Scan(&article.Author, &article.Id, &article.Title, &article.Content, &article.Date)
		if err != nil {
			log.Fatal(err)
		}
		articles = append(articles, article)
	}

	return articles, err
}

func (a ArticleModel) Update(db *sql.DB) error {
	return nil
}

func (a ArticleModel) Delete(db *sql.DB) error {
	return nil
}
