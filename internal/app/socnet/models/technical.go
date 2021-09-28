package models

import "time"

type GeneralTechFields struct {
	UUID      string
	Version   int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
