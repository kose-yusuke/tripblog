package model

import (
	"time"
	//"github.com/google/uuid"
)

// type Period struct {
// 	IsAllDay bool
// 	Begin    time.Time
// 	End      time.Time
// }

//var AllDay = Period{IsAllDay: true}


type Plan struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	UserID	   string `json:"userid"`
	// Content    string
}

func NewPlan(id string,title string, start time.Time, end time.Time,userid string) Plan {
	return Plan{
		Id:         id,
		Title:      title,
		Start: 		start,
		End:      	end,
		UserID: 	userid,
		// Content:    content,
	}
}