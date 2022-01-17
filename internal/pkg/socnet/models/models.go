package models

import "time"

type Session struct {
	UUID       string
	UserUUID   string
	ExpiringAt time.Time
	CreatedAt  time.Time
}

type User struct {
	GeneralTechFields
	FirstName    string
	LastName     string
	Bio          string
	Age          int32
	Gender       string
	Interests    string
	BirthDate    *time.Time
	Email        string
	PasswordHash string
}

type Friendship struct {
	GeneralTechFields
	RequesterUUID string
	ReceiverUUID  string
	Accepted      bool
}
