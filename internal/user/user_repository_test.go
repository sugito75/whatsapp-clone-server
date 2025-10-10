package user_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/sugito75/chat-app-server/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
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

func TestCreateUser_Success(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := user.NewRepository(gormDB)

	newUser := user.User{
		Username:   "John Doe",
		Phone:      "90121",
		LastOnline: time.Now(),
	}

	// Mock expected behavior
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
		WithArgs("John Doe", "90121", "", "", "", false, newUser.LastOnline).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	id, err := repo.CreateUser(newUser)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUser_Failure(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := user.NewRepository(gormDB)

	newUser := user.User{
		Username:   "John Doe",
		Phone:      "90121",
		LastOnline: time.Now(),
	}

	// Mock failure on insert
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
		WillReturnError(assert.AnError)
	mock.ExpectRollback()

	id, err := repo.CreateUser(newUser)

	assert.Error(t, err)
	assert.Equal(t, uint(0), id)
	assert.NoError(t, mock.ExpectationsWereMet())
}
