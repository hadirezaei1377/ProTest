package postgres_test

import (
	"ProTest/integration_test/repository/postgres"
	"ProTest/model"
	"os"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	pq "gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// all functionalities
func setup(t *testing.T) (*postgres.Database, *gorm.DB) {
	// connect to postgres database
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(pq.Open(dsn), &gorm.Config{})
	assert.Nil(t, err)

	// do migrations
	err := db.AutoMigrate(&model.User{})
	assert.Nil(t, err)

	// create and return gorm connection object as output
	return postgres.New(db), db

}
