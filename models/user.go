package models

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type UserModel struct {
	Id       int
	Name     []byte
	Email    []byte
	Password []byte
	Role     int
}

func (u UserModel) Create(db *sql.DB, email string, password []byte) (UserModel, error) {

	_, err := db.Exec("INSERT INTO C_USER (email, password, role) VALUES (?, ?, 1)", email, password)
	if err != nil {
		log.Fatal(err)
	}

	user, err := u.RetrieveByEmail(db, email)
	if err != nil {
		log.Fatal(err)
	}

	return user, err
}

func (u UserModel) RetrieveAll(db *sql.DB) ([]UserModel, error) {
	var users []UserModel

	rows, err := db.Query("SELECT uid, display_name, email, role FROM C_USER")
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

	var user UserModel
	err := db.QueryRow("SELECT uid, display_name, email, password, role FROM C_USER WHERE uid = ?", id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Email is %s\n", user.Email)
	}

	return user, err
}

func (u UserModel) RetrieveByEmail(db *sql.DB, email string) (UserModel, error) {

	var user UserModel
	err := db.QueryRow("SELECT uid, display_name, email, password, role FROM C_USER WHERE email = ?", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Username is %s\n", user.Email)
	}

	return user, err
}
