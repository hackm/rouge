package model

// Tag struct
type Tag struct {
	ID    uint32 `json:"id" form:"id" db:"id"`
	Name  string `json:"name" form:"name" db:"name"`
	Paper *Paper
}
