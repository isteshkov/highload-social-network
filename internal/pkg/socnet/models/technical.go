package models

import "time"

const ResponseBufferKey = "response_buffer"

type GeneralTechFields struct {
	UUID      string
	Version   int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}


type ResponseBuffer struct {
	Response   interface{}
	StatusCode int
}