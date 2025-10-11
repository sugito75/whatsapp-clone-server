package mock

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	return gormDB, mock
}
