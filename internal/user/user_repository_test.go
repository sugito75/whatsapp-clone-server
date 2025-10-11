package user_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/sugito75/chat-app-server/internal/user"
	"github.com/sugito75/chat-app-server/pkg/mock"
	"gorm.io/gorm"
)

func TestRepoCreateUser(t *testing.T) {
	t.Run("should success", func(t *testing.T) {
		gormDB, mock := mock.SetupMockDB(t)
		repo := user.NewRepository(gormDB)

		newUser := user.User{
			Username:   "John Doe",
			Phone:      "90121",
			LastOnline: time.Now(),
		}

		// Mock expected behavior
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "users"`).
			WithArgs("John Doe", "90121", "", nil, "", false, newUser.LastOnline).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		id, err := repo.CreateUser(newUser)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should fail when query is wrong", func(t *testing.T) {
		gormDB, mock := mock.SetupMockDB(t)
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
	})

}

func TestRepoGetUserByPhone(t *testing.T) {
	t.Run("should success", func(t *testing.T) {
		gormDB, mock := mock.SetupMockDB(t)
		repo := user.NewRepository(gormDB)

		expectedUser := user.User{
			ID:             1,
			Username:       "johndoe",
			Phone:          "08123456789",
			Password:       "hashedpassword",
			ProfilePicture: nil,
			Bio:            "Hello world",
			IsOnline:       true,
			LastOnline:     time.Now(),
		}

		rows := sqlmock.NewRows([]string{
			"id",
			"username",
			"phone",
			"password",
			"profile_picture",
			"bio",
			"is_online",
			"last_online",
		}).AddRow(
			expectedUser.ID,
			expectedUser.Username,
			expectedUser.Phone,
			expectedUser.Password,
			expectedUser.ProfilePicture,
			expectedUser.Bio,
			expectedUser.IsOnline,
			expectedUser.LastOnline,
		)

		mock.ExpectQuery(`SELECT .* FROM "users" WHERE phone = \$1 ORDER BY "users"\."id" LIMIT \$2`).
			WithArgs(expectedUser.Phone, 1).
			WillReturnRows(rows)

		u := repo.GetUserByPhone(expectedUser.Phone)

		assert.NotNil(t, u)
		assert.Equal(t, expectedUser.ID, u.ID)
		assert.Equal(t, expectedUser.Username, u.Username)
		assert.Equal(t, expectedUser.Phone, u.Phone)
		assert.Equal(t, expectedUser.Bio, u.Bio)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should fail when data is not found", func(t *testing.T) {

		gormDB, mock := mock.SetupMockDB(t)
		repo := user.NewRepository(gormDB)

		mock.ExpectQuery(`SELECT .* FROM "users" WHERE phone = \$1 ORDER BY "users"\."id" LIMIT \$2`).
			WithArgs("0999999999", 1).
			WillReturnError(gorm.ErrRecordNotFound)

		u := repo.GetUserByPhone("0999999999")

		assert.Nil(t, u)
		assert.NoError(t, mock.ExpectationsWereMet())

	})
}
