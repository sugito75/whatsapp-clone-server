package session

type SessionService interface {
	SaveSession(uid uint) error
	GetSessionID(uid uint) (string, error)
}
