package service

import (
    "github.com/gin-gonic/gin"
    "github.com/kose-yusuke/gocrud/api/cmd/db"
    "github.com/kose-yusuke/gocrud/api/cmd/calendar/model"
    //"time"
)

type UserRepository struct{}

//多分ここが何かおかしい
func (_ UserRepository) GetAll() ([]model.User, error) {
    db := db.GetDB()
    var u []model.User
    //if err := db.Table("users").Select("id").Scan(&u).Error; err != nil {
    if err := db.Find(&u).Error; err != nil {
        return nil, err
    }
    return u, nil
}

func (_ UserRepository) CreateModel(c *gin.Context) (model.User, error) {
    db := db.GetDB()
    var u model.User
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
    var u model.User

    if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
        return err
    }

    return nil
}