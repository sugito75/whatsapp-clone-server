package session_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/sugito75/chat-app-server/pkg/mock"
	"gorm.io/gorm"

	"github.com/sugito75/chat-app-server/pkg/session"
)

func TestSaveSession_Success(t *testing.T) {
	gormDB, mock := mock.SetupMockDB(t)
	svc := session.NewSessionService(gormDB)

	uid := uint(42)

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "sessions"`).
		WithArgs(uid, sqlmock.AnyArg()). // any id, user_id, session_id
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	err := svc.SaveSession(uid)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveSession_Failure(t *testing.T) {
	gormDB, mock := mock.SetupMockDB(t)
	svc := session.NewSessionService(gormDB)

	uid := uint(99)

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "sessions"`).
		WithArgs(uid, sqlmock.AnyArg()).
		WillReturnError(assert.AnError)
	mock.ExpectRollback()

	err := svc.SaveSession(uid)

	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetSessionID_Success(t *testing.T) {
	gormDB, mock := mock.SetupMockDB(t)
	svc := session.NewSessionService(gormDB)

	expectedSession := session.Session{
		UserID:    42,
		SessionID: "abc123xyz",
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "session_id"}).
		AddRow(1, expectedSession.UserID, expectedSession.SessionID)

	mock.ExpectQuery(`SELECT .* FROM "sessions" WHERE user_id = \$1 ORDER BY "sessions"\."id" LIMIT \$2`).
		WithArgs(expectedSession.UserID, 1).
		WillReturnRows(rows)

	sid, err := svc.GetSessionID(expectedSession.UserID)

	assert.NoError(t, err)
	assert.Equal(t, expectedSession.SessionID, sid)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetSessionID_NotFound(t *testing.T) {
	gormDB, mock := mock.SetupMockDB(t)
	svc := session.NewSessionService(gormDB)

	mock.ExpectQuery(`SELECT .* FROM "sessions" WHERE user_id = \$1 ORDER BY "sessions"\."id" LIMIT \$2`).
		WithArgs(uint(777), 1).
		WillReturnError(gorm.ErrRecordNotFound)

	sid, err := svc.GetSessionID(777)

	assert.Error(t, err)
	assert.Equal(t, "", sid)
	assert.NoError(t, mock.ExpectationsWereMet())
}
