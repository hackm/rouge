package content

import (
	"time"
)

// Content struct
type Content struct {
	ID        int64     `json:"id" form:"name"`
	PaperID   int64     `json:"paperId" form:"paperId"`
	Title     string    `json:"title" form:"title"`
	Body      []byte    `json:"body" form:"body"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
}
