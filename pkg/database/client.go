package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// required for stuff
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	client struct {
		*gorm.DB // Extends Gorm DB Struct to add our methods
	}

	// Client exports database functionalities
	Client interface {
		Ban(string, string) error
		Unban(string) error
		RegisterUser(string, string, string, string) (*User, error)
		GetAllUsers() ([]*User, error)
		GetUser(string) (*User, error)
		GetUsersByIP(string) ([]*User, error)
		DeleteUser(string) error
		Close() error
	}
)

func connect(username, password, host, dbname string) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, dbname)
	db, err := gorm.Open("mysql", connectionString)
	return db, err
}

func (c *client) Init() error {
	c.AutoMigrate(&User{})

	return nil
}

func (c *client) Close() error {
	return c.DB.Close()
}

// New Returns new DB Client or error
func New(username, password, host, dbname string) (Client, error) {
	db, err := connect(username, password, host, dbname)
	if err != nil {
		return nil, err
	}

	c := &client{db}

	if err = c.Init(); err != nil {
		return nil, err
	}

	return c, err
}
