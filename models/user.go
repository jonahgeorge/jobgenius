package models

import (
	"database/sql"
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type UserFactory struct {
}

type UserModel struct {
	Id          *int
	FirstName   *string
	LastName    *string
	DisplayName *string
	Email       *string
	Password    *string
	Role        *int
}

// Inserts a user account into db via email and password.
// The rest of the account information must be updated later.
// Called by UserController when using the Sign Up Form
func (u UserModel) Create(db *sql.DB, email string, password string) UserModel {

	_, err := db.Exec(
		"INSERT INTO Users (email, password, role) VALUES (?, ?, 1)", email, password)

	if err != nil {
		log.Println(err)
	}

	account := u.RetrieveByEmail(db, email)
	if err != nil {
		log.Println(err)
	}

	return account
}

// Retrieves all user accounts from the db
// Used by Account Controller for administrative user management
func (u UserModel) RetrieveAll(db *sql.DB) []UserModel {
	var users []UserModel

	rows, err := db.Query("SELECT uid, display_name, email, role FROM Users")
	if err != nil {
		log.Println("In UserModel->RetrieveAll")
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user UserModel
		err = rows.Scan(&user.Id, &user.DisplayName, &user.Email, &user.Role)
		if err != nil {
			log.Println("In UserModel->RetrieveAll")
			log.Println(err)
		}
		users = append(users, user)
	}

	return users
}

// Retrieves a single user account via user id (primary key)
func (u UserModel) RetrieveById(db *sql.DB, id string) UserModel {

	var user UserModel

	row := db.QueryRow("SELECT uid, display_name, email, password, role FROM Users WHERE uid = ?", id)
	err := row.Scan(&user.Id, &user.DisplayName, &user.Email, &user.Password, &user.Role)

	switch {
	case err != nil:
		log.Println(err)
	default:
		log.Printf("User record found for %s\n", user.Id)
	}

	return user
}

// Retrieves a single user account via user email
func (u UserModel) RetrieveByEmail(db *sql.DB, email string) UserModel {

	sql := `
	SELECT 
		uid, 
		display_name, 
		email, 
		password, 
		role 
	FROM
		Users 
	WHERE 
		email = ?`

	var user UserModel
	row := db.QueryRow(sql, email)

	err := row.Scan(
		&user.Id, &user.DisplayName, &user.Email,
		&user.Password, &user.Role)

	switch {
	case err != nil:
		log.Println(err)
	default:
		log.Printf("User record found for %s\n", user.Email)
	}

	return user
}
