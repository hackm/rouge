package model

import (
	"time"
)

// Paper struct
type Paper struct {
	ID        int64 `json:"id" form:"id" db:"id"`
	User      *User
	Content   *Content
	Tags      []*Tag
	CreatedAt time.Time `json:"createdAt" form:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" db:"updated_at"`
}
