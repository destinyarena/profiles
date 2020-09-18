package database

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User is the DB Structure for a user profile
type User struct {
	gorm.Model
	UUID    uuid.UUID `gorm:"primary_key"`
	Discord string    `gorm:"unique"`
	Faceit  string    `gorm:"unique"`
	Bungie  string    `gorm:"unique"`
	Banned  bool      `gorm:"default:false"`
	IPHash  string
}

// TableName sets the table name
func (u *User) TableName() string {
	return "profiles"
}

// BeforeCreate sets the user UUID
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()

	return scope.SetColumn("UUID", uuid)
}
