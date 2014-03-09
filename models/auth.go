package models

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	//	"log"
)

type AuthModel struct {
	Email    []byte
	Password []byte
}

func (a AuthModel) RetrieveByEmail(db *sql.DB, email []byte) (AuthModel, error) {
	var user AuthModel
	user.Email = email
	user.Password = []byte("Chase0the0Wolf")
	return user, nil
}
