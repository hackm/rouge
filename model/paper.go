package model

import (
	"time"
)

// Paper struct
type Paper struct {
	ID        uint32 `json:"id" form:"id" db:"id"`
	User      *User
	Content   *Content
	Tag       *Tag
	CreatedAt time.Time `json:"createdAt" form:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" db:"updated_at"`
}
