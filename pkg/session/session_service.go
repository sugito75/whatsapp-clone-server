package session

import (
	"crypto/rand"
	"encoding/base64"
	"log/slog"

	"gorm.io/gorm"
)

type sessionService struct {
	db *gorm.DB
}

func generateSessionID() string {
	b := make([]byte, 12)
	rand.Read(b)

	return base64.RawURLEncoding.EncodeToString(b)
}

func NewSessionService(db *gorm.DB) SessionService {
	return &sessionService{
		db: db,
	}
}

func (s *sessionService) SaveSession(uid uint) error {
	sessionID := generateSessionID()

	session := Session{
		UserID:    uid,
		SessionID: sessionID,
	}

	res := s.db.Create(&session)
	if res.Error != nil {
		slog.Error(res.Error.Error())
	}

	return res.Error
}

func (s *sessionService) GetSessionID(uid uint) (string, error) {
	var session Session
	result := s.db.First(&session, "user_id = $1", uid)

	if result.Error != nil {
		slog.Error(result.Error.Error())
		return "", result.Error
	}

	return session.SessionID, nil
}
