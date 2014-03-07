package controllers

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	. "github.com/jonahgeorge/husker/models"
	//	"html/template"
	"log"
	"net/http"
)

type ArticleController struct{}

func (a ArticleController) Index(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := ArticleModel{}.RetrieveAll(db)
		if err != nil {
			log.Fatal(err)
		}

		data := struct {
			Title    string
			Articles []ArticleModel
		}{
			"Articles",
			articles,
		}

		err = t.ExecuteTemplate(w, "articleIndex", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (a ArticleController) Retrieve(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		article, err := ArticleModel{}.RetrieveOne(db, r.FormValue("q"))
		if err != nil {
			log.Fatal(err)
		}

		data := struct {
			Title   string
			Article ArticleModel
		}{
			article.Title,
			article,
		}

		err = t.ExecuteTemplate(w, "articleShow", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (a ArticleController) Form(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "form", nil)
	}
}
