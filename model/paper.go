package model

import (
	"time"
)

// Paper struct
type Paper struct {
	ID        uint32
	User      User
	Content   Content
	CreatedAt time.Time
	UpdatedAt time.Time
}
