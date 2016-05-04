package model

import (
	"time"
)

// Stock struct
type Stock struct {
	ID        uint32    `json:"id" form:"id" db:"id"`
	UserID    uint32    `json:"userId" form:"userId" db:"user_id"`
	PaperID   uint32    `json:"paperId" form:"paperId" db:"paper_id"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt" db:"created_at"`
}
