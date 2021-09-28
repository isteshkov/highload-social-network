package models

import "time"

type User struct {
	UUID      string
	FirstName string
	LastName  string
	Bio       string
	Age       int32
	Gender    string
	Interests string
	BirthDate time.Time
}
