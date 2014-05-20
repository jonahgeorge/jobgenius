package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	. "github.com/jonahgeorge/jobgenius.net/models"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/russross/blackfriday"
)

type ArticleController struct{}

// Handles the rendering of all articles to the index page
func (a ArticleController) Index(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var articles []ArticleModel
		var err error

		// Parse GET parameters
		r.ParseForm()

		// Get article categories and session
		filters, err := ArticleFactory{}.GetCategories(db)
		session, err := store.Get(r, "user")

		if len(r.Form["title"]) > 0 {
			log.Println("Executing search by article title.")
			articles = ArticleFactory{}.RetrieveByName(db, r.Form["title"][0])
		} else if len(r.Form["filter"]) > 0 {
			// If filter was passed, use filter function
			articles, err = ArticleFactory{}.Filter(db, r.Form["filter"])
		} else {
			// If not, use normal retrieve
			articles, err = ArticleFactory{}.RetrieveAll(db)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = t.ExecuteTemplate(w, "article/index", map[string]interface{}{
			"Title":    "Articles",
			"Articles": articles,
			"Session":  session,
			"Filters":  filters,
		})

		if err != nil {
			log.Println(err)
		}
	}
}

// Handles the retrieval and rendering of a single article by the 'id' parameter
func (a ArticleController) Retrieve(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse url parameters
		params := mux.Vars(r)

		article, err := ArticleFactory{}.RetrieveById(db, params["id"])
		session, err := store.Get(r, "user")
		isAuthor := false

		if session.Values["Id"] == *article.User.Id {
			isAuthor = true
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = t.ExecuteTemplate(w, "article/show",
			map[string]interface{}{
				"Article": article,
				"Markdown": template.HTML(
					blackfriday.MarkdownCommon([]byte(*article.Body))),
				"Session":  session,
				"IsAuthor": isAuthor,
			})

		if err != nil {
			log.Println(err)
		}
	}
}

// Handles the rendering of the new article form page
func (a ArticleController) Form(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "user")
		filters, err := ArticleFactory{}.GetCategories(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.ExecuteTemplate(w, "article/form",
			map[string]interface{}{
				"Title":   "New Article",
				"Session": session,
				"Filters": filters,
			})
	}
}

// Handles the creation of articles from the article form page
func (a ArticleController) Create(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		id, err := ArticleModel{}.Create(db,
			map[string]interface{}{
				"uid":   session.Values["Id"],
				"title": r.FormValue("title"),
				"slug":  r.FormValue("slug"),
				"body":  r.FormValue("body"),
			})

		if err != nil {
			log.Printf("%s", err)
			http.Redirect(w, r, "/articles", http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/articles/%d", id), http.StatusTemporaryRedirect)
	}
}
