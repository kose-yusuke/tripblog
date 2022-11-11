package api

type Post struct {
    id      string  
    title string
    start time.Time 
    end time.Time 
    // UserID  uint   `gorm:"not null"  json:"user_id"`
}
