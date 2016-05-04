package model

import (
	"time"
)

// Content struct
type Content struct {
	ID        uint32    `json:"id" form:"name" db:"id"`
	PaperID   uint32    `json:"paperId" form:"paperId" db:"paper_id"`
	Title     string    `json:"title" form:"title" db:"title"`
	Body      []byte    `json:"body" form:"body" db:"body"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" db:"updated_at"`
}
