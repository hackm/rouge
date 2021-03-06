package entity

// User struct
type User struct {
	ID          int64  `json:"id" form:"id" db:"id"`
	Name        string `json:"name" form:"name" db:"name"`
	DisplayName string `json:"displayName" form:"displayName" db:"display_name"`
	Email       string `json:"email" form:"email" db:"email"`
}
