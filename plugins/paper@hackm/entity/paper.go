package entity

import (
  "time"
)

type Paper struct {
  Id int64 `json:"id" form:"id" db:"id"`
  UserId int64 
  ContentId int64
  // User *User
  Content *Content
  Tags []*Tag
  PublishedAt time.Time
  CreatedAt time.Time `json:"createdAt" form:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" db:"updated_at"`
}