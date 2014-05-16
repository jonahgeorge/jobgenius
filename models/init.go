package models

import (
	"encoding/gob"
)

func init() {
	gob.Register(UserModel{})
}
