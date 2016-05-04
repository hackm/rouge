package model

import (
	"time"
)

// Content struct
type Content struct {
	ID        uint32
	PaperID   uint32
	Title     string
	Body      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}
