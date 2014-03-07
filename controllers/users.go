package controllers

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	. "github.com/jonahgeorge/husker/models"
	"log"
	"net/http"
)

type UserController struct{}

func (u UserController) Index(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := UserModel{}.RetrieveAll(db)
		if err != nil {
			log.Fatal(err)
		}

		data := struct {
			Title string
			Users []UserModel
		}{
			"Users",
			users,
		}

		err = t.ExecuteTemplate(w, "userIndex", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (u UserController) Retrieve(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := UserModel{}.RetrieveById(db, r.FormValue("q"))
		articles, err := ArticleModel{}.RetrieveByAuthor(db, user.Id)
		interviews, err := InterviewModel{}.RetrieveByAuthor(db, user.Id)

		data := struct {
			Title      string
			User       UserModel
			Articles   []ArticleModel
			Interviews []InterviewModel
		}{
			"Users",
			user,
			articles,
			interviews,
		}

		err = t.ExecuteTemplate(w, "userShow", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
