package models

import (
	"log"
	"strings"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jonahgeorge/jobgenius.net/models/blocks"
)

type ArticleFactory struct {
}

func (a ArticleFactory) GetCategories() ([]blocks.Field, error) {

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

func (a ArticleFactory) GetRecent() []ArticleModel {

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
		Articles.published = 1
	ORDER BY
		Articles.timestamp
	LIMIT 5`

	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var articles []ArticleModel
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
	return articles
}

// Retrieve all articles
func (a ArticleFactory) RetrieveAll() ([]ArticleModel, error) {
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
func (a ArticleFactory) GetArticle(id string) ArticleModel {

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

	if err != nil {
		log.Println(err)
	}

	return article
}

// Retrieve a slice of articles authored by the user id parameter
func (a ArticleFactory) RetrieveByAuthor(id int) ([]ArticleModel, error) {
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

func (a ArticleFactory) Filter(filters []string) ([]ArticleModel, error) {
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

func (a ArticleFactory) RetrieveByName(name string) []ArticleModel {
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
		Articles.title LIKE ?
	OR
		Articles.slug LIKE ?`

	var articles []ArticleModel
	rows, err := db.Query(sql, "%"+name+"%", "%"+name+"%")
	if err != nil {
		return articles
	}

	for rows.Next() {
		var article ArticleModel
		err := rows.Scan(
			&article.User.Id, &article.User.DisplayName, &article.Id,
			&article.Title, &article.Slug, &article.Body, &article.Date)
		if err != nil {
			return articles
		}
		articles = append(articles, article)
	}

	return articles
}

type ArticleModel struct {
	Id    int
	User  UserModel
	Date  *time.Time
	Title *string
	Slug  *string
	Body  *string
}

// Create an article
func (a ArticleModel) Create(data map[string]interface{}) int64 {
	sql := `
	INSERT INTO Articles (title, slug, uid, body, published) 
	VALUES (?, ?, ?, ?, 0)`

	result, err := db.Exec(
		sql, data["title"], data["slug"],
		data["uid"], data["body"])

	if err != nil {
		log.Println(err)
	}
	id, err := result.LastInsertId()
	return id
}

func (a ArticleModel) AddCategory(id int64, category string) error {
	sql := `INSERT INTO Articles_Categories (id, value) VALUES (?, ?)`
	_, err := db.Exec(sql, id, category)
	return err
}

func (a ArticleModel) Update() error {
	return nil
}

func (a ArticleModel) Publish() error {
	sql := `UPDATE Articles SET published = 1 WHERE aid = ?`
	_, err := db.Exec(sql, a.Id)
	return err
}

func (a ArticleModel) Delete() error {
	sql := `DELETE FROM Articles WHERE aid = ?`
	_, err := db.Exec(sql, a.Id)
	return err
}
