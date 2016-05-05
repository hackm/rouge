package model

// Tag struct
type Tag struct {
	ID     int64  `json:"id" form:"id" db:"id"`
	Name   string `json:"name" form:"name" db:"name"`
	Papers []*Paper
}
