package entity

import(
  "time"
)

type Content struct {
  Id int64
  PaperId int64
  Title string
  Body string
  CreatedAt time.Time `json:"createdAt" form:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" db:"updated_at"`
}