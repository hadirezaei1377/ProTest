package postgres_test

import (
	"ProTest/integration_test/repository/postgres"
	"ProTest/model"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

/* in this file we should write test for each function or method in our module and do some basic works like creating database and
and do migration(doing all models and relation that we need to our empty db like creating users table)
instead of testing each module here, its better to create a file named main_test and add all functionality that are giong to use there
*/

func TestGetUserByEmail(t *testing.T) {
	db, conn := setup(t)

	// create random data
	faker.Email()
	user := &model.User{
		ID:       0,
		Email:    faker.Email(),
		Password: faker.Password(),
		Name:     faker.Name(),
	}

	err := conn.Create(user).Error
	assert.Nil(t, err)

	return postgres.New(db), db
}
