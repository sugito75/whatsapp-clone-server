package user

import "time"

type User struct {
	ID             uint
	Username       string
	Phone          string
	Password       string
	ProfilePicture string
	Bio            string
	IsOnline       bool
	LastOnline     time.Time
}
