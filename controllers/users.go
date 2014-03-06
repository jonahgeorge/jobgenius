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
