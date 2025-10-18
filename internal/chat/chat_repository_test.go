package chat_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/sugito75/chat-app-server/internal/chat"
	"github.com/sugito75/chat-app-server/pkg/mock"
)

func TestRepoCreateChat(t *testing.T) {
	t.Run("should success", func(t *testing.T) {
		db, mock := mock.SetupMockDB(t)
		repo := chat.NewRepo(db)

		now := time.Now()
		title := "Test Chat"
		desc := "Description"
		icon := "icon.png"

		c := chat.Chat{
			ChatType:    "group",
			Title:       &title,
			Description: &desc,
			Icon:        &icon,
			CreatedBy:   1,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		m := chat.Message{}

		// Expect INSERT query
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(
			`INSERT INTO "chats" ("chat_type","title","description","icon","created_by","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
			WithArgs(c.ChatType, c.Title, c.Description, c.Icon, c.CreatedBy, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(10))
		mock.ExpectCommit()

		id, err := repo.CreateChat(c, m)

		assert.NoError(t, err)
		assert.Equal(t, uint64(10), id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should fail", func(t *testing.T) {
		db, mock := mock.SetupMockDB(t)
		repo := chat.NewRepo(db)

		title := "Bad Chat"
		c := chat.Chat{
			ChatType:  "private",
			Title:     &title,
			CreatedBy: 99,
		}

		m := chat.Message{}

		// Expect failure on insert
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(
			`INSERT INTO "chats" ("chat_type","title","description","icon","created_by","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
			WithArgs(c.ChatType, c.Title, c.Description, c.Icon, c.CreatedBy, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(assert.AnError)
		mock.ExpectRollback()

		id, err := repo.CreateChat(c, m)

		assert.Error(t, err)
		assert.Equal(t, uint64(0), id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
