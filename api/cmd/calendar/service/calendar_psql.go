package service

import (
    "github.com/gin-gonic/gin"
    "github.com/kose-yusuke/gocrud/api/cmd/db"
    "github.com/kose-yusuke/gocrud/api/cmd/calendar/model"
)

type CalRepository struct{}


type Calendar model.Calendar

type CalendarDetail struct {
    ID     string
	UserID string
	Name   string
	Color  model.Color
	Plans  []Plan
	Shares []string
}


func (_ CalRepository) GetAll() ([]CalendarDetail, error) {
    db := db.GetDB()
    var u []CalendarDetail
    if err := db.Table("calendar").Select("id").Scan(&u).Error; err != nil {
        return nil, err
    }
    return u, nil
}


func (_ CalRepository) CreateModel(c *gin.Context) (Calendar, error) {
    db := db.GetDB()
    var u Calendar
    if err := c.BindJSON(&u); err != nil {
        return u, err
    }
    if err := db.Create(&u).Error; err != nil {
        return u, err
    }
    return u, nil
}


func (_ CalRepository) GetByID(id string) (model.Calendar, error) {
    db := db.GetDB()
    var u model.Calendar
    if err := db.Where("id = ?", id).First(&u).Error; err != nil {
        return u, err
    }

    return u, nil
}


func (_ CalRepository) UpdateByID(id string, c *gin.Context) (model.Calendar, error) {
    db := db.GetDB()
    var u model.Calendar
    if err := db.Where("id = ?", id).First(&u).Error; err != nil {
        return u, err
    }
    if err := c.BindJSON(&u); err != nil {
        return u, err
    }
    db.Save(&u)

    return u, nil
}


func (_ CalRepository) DeleteByID(id string) error {
    db := db.GetDB()
    var u Calendar

    if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
        return err
    }

    return nil
}