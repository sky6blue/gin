package user

import (
	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID
}

func New() User {
	return User{
		Id: uuid.New(),
	}
}
