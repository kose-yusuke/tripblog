package model

import (
	"time"
	"github.com/google/uuid"
)

// type Period struct {
// 	IsAllDay bool
// 	Begin    time.Time
// 	End      time.Time
// }

//var AllDay = Period{IsAllDay: true}


type Plan struct {
	ID         string
	CalendarID string
	UserID     string
	Name       string
	Memo       string
	Color      Color
	Private    bool
	Shares     []string
	IsAllDay bool
	Begin    time.Time
	End      time.Time
}

func NewPlan(calendarID, userID, name, memo string, color Color, private bool, shares []string, isallday bool, begin, end time.Time) Plan {
	return Plan{
		ID:         uuid.New().String(),
		CalendarID: calendarID,
		UserID:     userID,
		Name:       name,
		Memo:       memo,
		Color:      color,
		Private:    private,
		Shares:     shares,
		IsAllDay:   isallday,
		Begin: 		begin,
		End:      	end,
	}
}