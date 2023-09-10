package postgres

import (
	"ProTest/model"

	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
}

// get connection from postgres
func New(connection *gorm.DB) *Database {
	return &Database{connection: connection}
}

// define some methods for our object

// functionality for module which connected to db
func (d *Database) GetUserByEmail(email string) (*model.User, error) {
	user := new(model.User)
	if err := d.connection.Model(user).Where("email = ?", email).
		First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
