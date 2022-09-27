package model

import (
	//"github.com/google/uuid"
)

type Calendar struct {
	User_id     string
	Plan_id 	string
}

func NewCalendar(user_id string, plan_id string) Calendar {
	return Calendar{
		User_id:     user_id,
		Plan_id:     plan_id,
	}
}