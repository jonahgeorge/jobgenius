package models

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type AccountModel struct {
	Id       int
	Name     []byte
	Email    []byte
	Password []byte
	Role     int
}

func (a AccountModel) Create(db *sql.DB, email string, password []byte) AccountModel {
	_, err := db.Exec("INSERT INTO C_USER (email, password, role) VALUES (?, ?, 1)", email, password)
	if err != nil {
		log.Fatal(err)
	}

	account := a.RetrieveByEmail(db, email)
	return account
}

func (a AccountModel) RetrieveAll(db *sql.DB) ([]AccountModel, error) {
	var accounts []AccountModel

	rows, err := db.Query("SELECT uid, display_name, email, role FROM C_USER")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var account AccountModel
		if err = rows.Scan(&account.Id, &account.Name, &account.Email, &account.Role); err != nil {
			log.Fatal(err)
		}
		accounts = append(accounts, account)
	}
	return accounts, err
}

func (a AccountModel) RetrieveById(db *sql.DB, id string) (AccountModel, error) {
	var account AccountModel
	err := db.QueryRow("SELECT uid, display_name, email, password, role FROM C_USER WHERE uid = ?", id).Scan(&account.Id, &account.Name, &account.Email, &account.Password, &account.Role)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Account found for %s\n", account.Id)
	}

	return account, err
}

func (a AccountModel) RetrieveByEmail(db *sql.DB, email string) AccountModel {

	var account AccountModel
	err := db.QueryRow("SELECT uid, display_name, email, password, role FROM C_USER WHERE email = ?", email).Scan(&account.Id, &account.Name, &account.Email, &account.Password, &account.Role)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No account with that email.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Account found for %s\n", account.Email)
	}

	return account
}
