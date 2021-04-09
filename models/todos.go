package models

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// UUID contains id with type uuid
type UUID struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
}

// BeforeCreate Add uuid value into ID column
func (u *UUID) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID, err = uuid.NewV4()
	if err != nil {
		err = errors.New("can't save invalid data")
	}
	return
}

// Todos godoc
type Todos struct {
	UUID
	Name    string
	Done    bool
	Deleted bool
}
