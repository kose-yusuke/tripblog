package model

import (
	"time"
)

// Userの定義
type User struct {
	ID 			string
	Name 		string
	Created_at 	time.Time
	Updated_at  time.Time
}

func NewUser(id string, name string, created_at time.Time, updated_at time.Time) User {
	return User{
		ID: id,
		Name: name,
		Created_at: created_at,
		Updated_at: updated_at,
	}
}

// type User struct {
//     ID    uint   `json:"id" binding:"required"`
//     Name  string `json:"name" binding:"required"`
//     Posts []Post `json:"posts"`
// }