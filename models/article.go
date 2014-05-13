package models

import (
	"database/sql"
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type ArticleFactory struct {
}

type ArticleModel struct {
	Id      *int
	Author  *string
	Picture *string
	Date    *string
	Title   *string
	Content *string
}

// Create an article
func (a ArticleModel) Create(db *sql.DB, data map[string]interface{}) (int64, error) {

	sql := `
	INSERT INTO 
	C_ARTICLE (title, uid, body, published) 
	VALUES (?, ?, ?, 1)`

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

	sql := `
	SELECT 
		C_USER.display_name
	,	C_USER.email_hash
	,	C_ARTICLE.aid
	,	C_ARTICLE.title
	,	C_ARTICLE.body
	FROM 
		C_ARTICLE
	 LEFT JOIN 
		C_USER ON C_ARTICLE.uid = C_USER.uid
	WHERE 
		C_ARTICLE.published = 1`

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var article ArticleModel

		err = rows.Scan(&article.Author, &article.Picture, &article.Id, &article.Title, &article.Content)
		if err != nil {
			log.Fatal(err)
		}

		articles = append(articles, article)
	}
	return articles, err
}

// Retrieve one article by article id (primary key)
func (a ArticleModel) RetrieveById(db *sql.DB, id string) (ArticleModel, error) {

	sql := `
	SELECT 
		C_USER.display_name
	,	C_USER.email_hash
	,	C_ARTICLE.aid
	,	C_ARTICLE.title
	,	C_ARTICLE.body 
	,	C_ARTICLE.timestamp
	FROM 
		C_ARTICLE
	LEFT JOIN 
		C_USER ON C_ARTICLE.uid = C_USER.uid
	WHERE 
		C_ARTICLE.published = 1 
	 AND 
		C_ARTICLE.aid = ?`

	var article ArticleModel
	row := db.QueryRow(sql, id)
	err := row.Scan(&article.Author, &article.Picture, &article.Id,
		&article.Title, &article.Content, &article.Date)
	return article, err
}

// Retrieve a slice of articles authored by the user id parameter
func (a ArticleModel) RetrieveByAuthor(db *sql.DB, id int) ([]ArticleModel, error) {
	var articles []ArticleModel

	sql := `
	SELECT 
		C_USER.display_name
	,	C_USER.email_hash
	,	C_ARTICLE.aid
	,	C_ARTICLE.title
	,	C_ARTICLE.body
	,	C_ARTICLE.timestamp
	FROM 
		C_ARTICLE
	LEFT JOIN 
		C_USER ON C_ARTICLE.uid = C_USER.uid
	WHERE 
		C_ARTICLE.published = 1 
	 AND 
		C_USER.uid = ?`

	rows, err := db.Query(sql, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var article ArticleModel
		err = rows.Scan(&article.Author, &article.Picture, &article.Id, &article.Title, &article.Content, &article.Date)
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
