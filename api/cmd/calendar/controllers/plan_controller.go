package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/kose-yusuke/gocrud/api/cmd/calendar/service"
)


type PlanController struct{}

// Index action: GET /users
func (pc PlanController) Index(c *gin.Context) {
    var u service.PlanRepository
    p, err := u.GetAll()
    if err != nil {
        c.AbortWithStatus(404)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    } else {
        c.JSON(200, p)
    }
}

// Create action: POST /users
func (pc PlanController) Create(c *gin.Context) {
    var u service.PlanRepository
    p, err := u.CreateModel(c)

    if err != nil {
        c.AbortWithStatus(400)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    } else {
        c.JSON(201, p)
    }
}

// Show action: Get /users/:id
func (pc PlanController) Show(c *gin.Context) {
    id := c.Params.ByName("id")
    var u service.PlanRepository
    user, err := u.GetByID(id)

    if err != nil {
        c.AbortWithStatus(400)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    } else {
        c.JSON(200, user)
    }
}

// Update action: Put /users/:id
func (pc PlanController) Update(c *gin.Context) {
    id := c.Params.ByName("id")
    var u service.PlanRepository
    p, err := u.UpdateByID(id, c)

    if err != nil {
        c.AbortWithStatus(404)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    } else {
        c.JSON(200, p)
    }
}

// Delete action: DELETE /users/:id
func (pc PlanController) Delete(c *gin.Context) {
    id := c.Params.ByName("id")
    var u service.PlanRepository
    if err := u.DeleteByID(id); err != nil {
        c.AbortWithStatus(403)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"success": "ID" + id + "のユーザーを削除しました"})
    return
}