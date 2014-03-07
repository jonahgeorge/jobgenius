package models

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type UserModel struct {
	Id       int
	Name     string
	Email    string
	Password string
	Role     int
}

func (u UserModel) Create(db *sql.DB) error {
	return nil
}

func (u UserModel) RetrieveAll(db *sql.DB) ([]UserModel, error) {
	var users []UserModel

	sql := `SELECT uid, display_name, email, role FROM C_USER`

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u UserModel
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Role)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	return users, err
}

func (u UserModel) RetrieveById(db *sql.DB, id string) (UserModel, error) {

	sql := `SELECT uid, display_name, email, role FROM C_USER WHERE uid = ` + id

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user UserModel
	err = db.QueryRow(sql).Scan(&user.Id, &user.Name, &user.Email, &user.Role)
	if err != nil {
		log.Fatal(err)
	}

	return user, err
}
