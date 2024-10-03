package model

import (
	"github.com/google/uuid"
)

type Client struct {
	Id    *uuid.UUID `db:"id"`
	Name  string     `db:"name"`
	Email string     `db:"email"`
	Phone string     `db:"phone"`
}
