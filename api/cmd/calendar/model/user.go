package model

// Userの定義
type User struct {
	ID string
}

func NewUser(id string) User {
	return User{
		ID: id,
	}
}

// type User struct {
//     ID    uint   `json:"id" binding:"required"`
//     Name  string `json:"name" binding:"required"`
//     Posts []Post `json:"posts"`
// }