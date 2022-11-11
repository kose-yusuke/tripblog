package model

import (
	"time"
)

// Userの定義
type User struct {
	ID 			string `json:"id"`
	Name 		string	`json:"name"`
	Created_at 	time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	Plans []Plan `json:"plans"`

}

// func NewUser(id string, name string, created_at time.Time, updated_at time.Time,) User {
// 	return User{
// 		ID: id,
// 		Name: name,
// 		Created_at: created_at,
// 		Updated_at: updated_at,
// 		Plans: plans
// 	}
// }

// type User struct {
//     ID    uint   `json:"id" binding:"required"`
//     Name  string `json:"name" binding:"required"`
//     Posts []Post `json:"posts"`
// }