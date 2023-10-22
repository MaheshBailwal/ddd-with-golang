package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Name      string
	Age       int
	CreatedAt time.Time `dbfield:"Created_At"`
	IsDeleted bool
}
