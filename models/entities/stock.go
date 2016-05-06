package entity

import (
	"time"
)

// Stock struct
type Stock struct {
	ID        int64     `json:"id" form:"id" db:"id"`
	UserID    int64     `json:"userId" form:"userId" db:"user_id"`
	PaperID   int64     `json:"paperId" form:"paperId" db:"paper_id"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt" db:"created_at"`
}
