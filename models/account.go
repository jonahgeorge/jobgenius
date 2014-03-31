package models

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

type AccountModel struct {
	Id       sql.NullInt64
	Name     sql.NullString
	Email    sql.NullString
	Password sql.NullString
	Role     sql.NullInt64
}

// Inserts a user account into db via email and password.
// The rest of the account information must be updated later.
// Called by UserController when using the Sign Up Form
func (a AccountModel) Create(db *sql.DB, email string, password string) AccountModel {
	_, err := db.Exec("INSERT INTO C_USER (email, password, role) VALUES (?, ?, 1)", email, password)
	if err != nil {
		log.Printf("%s", err)
	}

	account := a.RetrieveByEmail(db, email)
	if err != nil {
		log.Printf("%s", err)
	}

	return account
}

// Retrieves all user accounts from the db
// Used by Account Controller for administrative user management
func (a AccountModel) RetrieveAll(db *sql.DB) []AccountModel {

	var accounts []AccountModel

	rows, err := db.Query("SELECT uid, display_name, email, role FROM C_USER")
	if err != nil {
		log.Printf("%s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var account AccountModel

		err = rows.Scan(&account.Id, &account.Name, &account.Email, &account.Role)
		if err != nil {
			log.Printf("%s", err)
		}

		accounts = append(accounts, account)
	}
	return accounts
}

// Retrieves a single user account via user id (primary key)
func (a AccountModel) RetrieveById(db *sql.DB, id string) AccountModel {

	var account AccountModel

	row := db.QueryRow("SELECT uid, display_name, email, password, role FROM C_USER WHERE uid = ?", id)
	err := row.Scan(&account.Id, &account.Name, &account.Email, &account.Password, &account.Role)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Printf("%s", err)
	default:
		log.Printf("Account found for %s\n", account.Id.Int64)
	}

	return account
}

// Retrieves a single user account via user email
func (a AccountModel) RetrieveByEmail(db *sql.DB, email string) AccountModel {

	var account AccountModel

	row := db.QueryRow("SELECT uid, display_name, email, password, role FROM C_USER WHERE email = ?", email)
	err := row.Scan(&account.Id, &account.Name, &account.Email, &account.Password, &account.Role)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No account with that email.")
	case err != nil:
		log.Printf("%s", err)
	default:
		log.Printf("Account found for %s\n", account.Email.String)
	}

	return account
}
