package models

import (
	"encoding/gob"
)

type UserModel struct {
	Email []byte
	Name  []byte
	Id    int
}

func init() {
	gob.Register(UserModel{})
}
