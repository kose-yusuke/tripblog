package service

import (
    "github.com/gin-gonic/gin"
    "github.com/kose-yusuke/gocrud/api/cmd/db"
    "github.com/kose-yusuke/gocrud/api/cmd/calendar/model"
)

type UserRepository struct{}


type User model.User

type UserProfile struct {
    Id   int
}


func (_ UserRepository) GetAll() ([]UserProfile, error) {
    db := db.GetDB()
    var u []UserProfile
    if err := db.Table("users").Select("id").Scan(&u).Error; err != nil {
        return nil, err
    }
    return u, nil
}


func (_ UserRepository) CreateModel(c *gin.Context) (User, error) {
    db := db.GetDB()
    var u User
    if err := c.BindJSON(&u); err != nil {
        return u, err
    }
    if err := db.Create(&u).Error; err != nil {
        return u, err
    }
    return u, nil
}


func (_ UserRepository) GetByID(id string) (model.User, error) {
    db := db.GetDB()
    var me model.User
    if err := db.Where("id = ?", id).First(&me).Error; err != nil {
        return me, err
    }

    return me, nil
}


func (_ UserRepository) UpdateByID(id string, c *gin.Context) (model.User, error) {
    db := db.GetDB()
    var u model.User
    if err := db.Where("id = ?", id).First(&u).Error; err != nil {
        return u, err
    }
    if err := c.BindJSON(&u); err != nil {
        return u, err
    }
    db.Save(&u)

    return u, nil
}


func (_ UserRepository) DeleteByID(id string) error {
    db := db.GetDB()
    var u User

    if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
        return err
    }

    return nil
}