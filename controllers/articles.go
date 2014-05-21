package controllers

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	. "github.com/jonahgeorge/jobgenius.net/models"

	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
)

type ArticleController struct{}

// Handles the rendering of all articles to the index page
func (a ArticleController) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var articles []ArticleModel
		var err error

		// Parse GET parameters
		r.ParseForm()

		// Get article categories and session
		categories, err := ArticleFactory{}.GetCategories()
		session, err := store.Get(r, "user")

		if len(r.Form["title"]) > 0 {
			log.Println("Executing search by article title.")
			articles = ArticleFactory{}.RetrieveByName(r.Form["title"][0])
		} else if len(r.Form["filter"]) > 0 {
			// If filter was passed, use filter function
			articles, err = ArticleFactory{}.Filter(r.Form["filter"])
		} else {
			// If not, use normal retrieve
			articles, err = ArticleFactory{}.RetrieveAll()
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.ExecuteTemplate(w, "article/index", map[string]interface{}{
			"Title":      "Articles",
			"Articles":   articles,
			"Session":    session,
			"Categories": categories,
		})

		if err != nil {
			log.Println(err)
		}
	}
}

// Handles the retrieval and rendering of a single article by the 'id' parameter
func (a ArticleController) Retrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse url parameters
		params := mux.Vars(r)

		article := ArticleFactory{}.GetArticle(params["id"])
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
				"Date":     article.Date.Format("January 2, 2006"),
			})

		if err != nil {
			log.Println(err)
		}
	}
}

// Handles the rendering of the new article form page
func (a ArticleController) Form() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "user")
		categories, err := ArticleFactory{}.GetCategories()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.ExecuteTemplate(w, "article/form",
			map[string]interface{}{
				"Title":      "New Article",
				"Session":    session,
				"Categories": categories,
			})
	}
}

// Handles the creation of articles from the article form page
func (a ArticleController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "user")

		// Get `Multiple` POST parameters
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		id := ArticleModel{}.Create(map[string]interface{}{
			"uid":   session.Values["Id"],
			"title": r.Form.Get("title"),
			"slug":  r.Form.Get("slug"),
			"body":  r.Form.Get("body"),
		})

		for _, category := range r.Form["category"] {
			ArticleModel{}.AddCategory(id, category)
		}

		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/articles", http.StatusFound)
	}
}

func (a ArticleController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		article := ArticleFactory{}.GetArticle(mux.Vars(r)["id"])
		if err := article.Delete(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/articles", http.StatusFound)
	}
}

func (a ArticleController) Publish() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		article := ArticleFactory{}.GetArticle(mux.Vars(r)["id"])
		if err := article.Publish(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/articles", http.StatusFound)
	}
}

func (a ArticleController) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
