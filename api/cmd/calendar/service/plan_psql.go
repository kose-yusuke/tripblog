package service

import (
	"time"
    "github.com/gin-gonic/gin"
    "github.com/kose-yusuke/gocrud/api/cmd/db"
    "github.com/kose-yusuke/gocrud/api/cmd/calendar/model"
)

type PlanRepository struct{}


type Plan model.Plan

type PlanDetail struct {
    ID         string
	Title      string
	Start      time.Time
	End        time.Time
    Content    string
}


func (_ PlanRepository) GetAll() ([]PlanDetail, error) {
    db := db.GetDB()
    var u []PlanDetail
    if err := db.Table("plans").Select("id").Scan(&u).Error; err != nil {
        return nil, err
    }
    return u, nil
}


func (_ PlanRepository) CreateModel(c *gin.Context) (Plan, error) {
    db := db.GetDB()
    var u Plan
    if err := c.BindJSON(&u); err != nil {
        return u, err
    }
    if err := db.Create(&u).Error; err != nil {
        return u, err
    }
    return u, nil
}


func (_ PlanRepository) GetByID(id string) (model.Plan, error) {
    db := db.GetDB()
    var u model.Plan
    if err := db.Where("id = ?", id).First(&u).Error; err != nil {
        return u, err
    }

    return u, nil
}


func (_ PlanRepository) UpdateByID(id string, c *gin.Context) (model.Plan, error) {
    db := db.GetDB()
    var u model.Plan
    if err := db.Where("id = ?", id).First(&u).Error; err != nil {
        return u, err
    }
    if err := c.BindJSON(&u); err != nil {
        return u, err
    }
    db.Save(&u)

    return u, nil
}


func (_ PlanRepository) DeleteByID(id string) error {
    db := db.GetDB()
    var u Plan

    if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
        return err
    }

    return nil
}