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
	ID         string
	Title      string
	Start      time.Time
	End        time.Time
	Content    string
}

func NewPlan(id string,title string, start time.Time, end time.Time, content string) Plan {
	return Plan{
		ID:         id,
		Title:      title,
		Start: 		start,
		End:      	end,
		Content:    content,
	}
}