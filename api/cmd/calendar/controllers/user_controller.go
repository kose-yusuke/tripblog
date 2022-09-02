package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/kose-yusuke/gocrud/api/cmd/calendar/service"
)


type UserController struct{}

// Index action: GET /users
func (pc UserController) Index(c *gin.Context) {
    var u service.UserRepository
    p, err := u.GetAll()
    if err != nil {
        c.AbortWithStatus(404)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    } else {
        c.JSON(200, p)
    }
}

// Create action: POST /users
func (pc UserController) Create(c *gin.Context) {
    var u service.UserRepository
    p, err := u.CreateModel(c)

    if err != nil {
        c.AbortWithStatus(400)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    } else {
        c.JSON(201, p)
    }
}

// Show action: Get /users/:id
func (pc UserController) Show(c *gin.Context) {
    id := c.Params.ByName("id")
    var u service.UserRepository
    user, err := u.GetByID(id)

    if err != nil {
        c.AbortWithStatus(400)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    } else {
        c.JSON(200, user)
    }
}

// Update action: Put /users/:id
func (pc UserController) Update(c *gin.Context) {
    id := c.Params.ByName("id")
    var u service.UserRepository
    p, err := u.UpdateByID(id, c)

    if err != nil {
        c.AbortWithStatus(404)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    } else {
        c.JSON(200, p)
    }
}

// Delete action: DELETE /users/:id
func (pc UserController) Delete(c *gin.Context) {
    id := c.Params.ByName("id")
    var u service.UserRepository
    if err := u.DeleteByID(id); err != nil {
        c.AbortWithStatus(403)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"success": "ID" + id + "のユーザーを削除しました"})
    return
}