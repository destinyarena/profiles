package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
    client struct {
        *gorm.DB // Extends Gorm DB Struct to add our methods
    }

    Client interface {
        Ban(string) error
        RegisterUser(string, string, string) (error, *User)
        GetAllUsers() (error, []User)
        GetUser(string) (error, *User)
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

/*
Returns new DB Client or error
*/
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

