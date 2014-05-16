package models

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jonahgeorge/jobgenius.net/models/blocks"
)

type ArticleFactory struct {
}

func (a ArticleFactory) GetCategories(db *sql.DB) ([]blocks.Field, error) {

	sql := `
	SELECT *
	FROM Articles_Categories_Lookup`

	var fields []blocks.Field

	rows, err := db.Query(sql)
	if err != nil {
		return fields, err
	}

	for rows.Next() {
		var field blocks.Field
		err := rows.Scan(&field.Key, &field.Value)
		if err != nil {
			return fields, err
		}
		fields = append(fields, field)
	}

	return fields, err
}

// Retrieve all articles
func (a ArticleFactory) RetrieveAll(db *sql.DB) ([]ArticleModel, error) {
	var articles []ArticleModel

	sql := `
	SELECT 
		Users.uid
	,	Users.display_name
	,	Articles.aid
	,	Articles.title
	,	Articles.slug
	,	Articles.body
	FROM 
		Articles
	 LEFT JOIN 
		Users ON Articles.uid = Users.uid
	WHERE 
		Articles.published = 1`

	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var article ArticleModel

		err = rows.Scan(
			&article.User.Id, &article.User.DisplayName,
			&article.Id, &article.Title, &article.Slug, &article.Body)

		if err != nil {
			log.Println(err)
		}

		articles = append(articles, article)
	}
	return articles, err
}

// Retrieve one article by article id (primary key)
func (a ArticleFactory) RetrieveById(db *sql.DB, id string) (ArticleModel, error) {

	sql := `
	SELECT 
		Users.uid
	,	Users.display_name
	,	Articles.aid
	,	Articles.title
	,	Articles.slug
	,	Articles.body 
	,	Articles.timestamp
	FROM 
		Articles
	LEFT JOIN 
		Users ON Articles.uid = Users.uid
	WHERE 
		Articles.published = 1 
	 AND 
		Articles.aid = ?`

	var article ArticleModel
	row := db.QueryRow(sql, id)

	err := row.Scan(
		&article.User.Id, &article.User.DisplayName, &article.Id,
		&article.Title, &article.Slug, &article.Body, &article.Date)

	return article, err
}

// Retrieve a slice of articles authored by the user id parameter
func (a ArticleFactory) RetrieveByAuthor(db *sql.DB, id int) ([]ArticleModel, error) {
	var articles []ArticleModel

	sql := `
	SELECT 
		Users.uid
	,	Users.display_name
	,	Articles.aid
	,	Articles.title
	,	Articles.slug
	,	Articles.body
	,	Articles.timestamp
	FROM 
		Articles
	LEFT JOIN 
		Users ON Articles.uid = Users.uid
	WHERE 
		Articles.published = 1 
	 AND 
		Users.uid = ?`

	rows, err := db.Query(sql, id)
	if err != nil {
		return articles, err
	}
	defer rows.Close()

	for rows.Next() {
		var article ArticleModel
		err = rows.Scan(
			&article.User.Id, &article.User.DisplayName, &article.Id,
			&article.Title, &article.Slug, &article.Body, &article.Date)
		if err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}

	return articles, err
}

func (a ArticleFactory) Filter(db *sql.DB, filters []string) ([]ArticleModel, error) {
	sql := `
	SELECT 
		Users.uid
	,	Users.display_name
	,	Articles.aid
	,	Articles.title
	,	Articles.slug
	,	Articles.body 
	,	Articles.timestamp
	FROM 
		Articles
	LEFT JOIN 
		Users ON Articles.uid = Users.uid
	LEFT JOIN
		Articles_Categories ON Articles_Categories.id = Articles.aid
	LEFT JOIN
		Articles_Categories_Lookup ON Articles_Categories_Lookup.id = Articles_Categories.value
	WHERE 
		Articles.published = 1
	AND
		Articles_Categories_Lookup.value IN ('` + strings.Join(filters, "','") + `')
	GROUP BY
		Articles.aid`

	var articles []ArticleModel
	rows, err := db.Query(sql)
	if err != nil {
		return articles, err
	}

	for rows.Next() {
		var article ArticleModel
		err := rows.Scan(
			&article.User.Id, &article.User.DisplayName, &article.Id,
			&article.Title, &article.Slug, &article.Body, &article.Date)
		if err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}

	return articles, err
}

type ArticleModel struct {
	Id    *int
	User  UserModel
	Date  *string
	Title *string
	Slug  *string
	Body  *string
}

// Create an article
func (a ArticleModel) Create(db *sql.DB, data map[string]interface{}) (int64, error) {

	sql := `
	INSERT INTO 
	Articles (title, subtitle, uid, body, published) 
	VALUES (?, ?, ?, ?, 1)`

	result, err := db.Exec(sql, data["title"], data["slug"], data["uid"], data["body"])
	if err != nil {
		log.Printf("%s", err)
	}

	id, err := result.LastInsertId()

	return id, err
}

func (a ArticleModel) Update(db *sql.DB) error {
	return nil
}

func (a ArticleModel) Delete(db *sql.DB) error {
	return nil
}
