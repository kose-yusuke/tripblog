package server

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "github.com/kose-yusuke/tripblog/api/cmd/calendar/controllers"
    "time"
)

// Init is initialize server
func Init() {
    r := router()
    r.Run()
}

func router() *gin.Engine {
    r := gin.Default()

    setCors(r)

    u := r.Group("/users")
    {
        ctrl := controllers.UserController{}
		//第一引数にパス、第二引数に実行したい内容ハンドラ
        u.GET("", ctrl.Index)
        u.POST("", ctrl.Create)
        u.GET("/:id", ctrl.Show)
        u.PUT("/:id", ctrl.Update)
        u.DELETE("/:id", ctrl.Delete)
    }

    p := r.Group("/plans")
    {
        ctrl := controllers.PlanController{}
        p.GET("", ctrl.Index)
        p.POST("", ctrl.Create)
        p.GET("/:id", ctrl.Show)
        p.PUT("/:id", ctrl.Update)
        p.DELETE("/:id", ctrl.Delete)
    }

	c := r.Group("/calendar")
    {
        ctrl := controllers.CalController{}
        c.GET("", ctrl.Index)
        c.POST("", ctrl.Create)
        c.GET("/:id", ctrl.Show)
        c.PUT("/:id", ctrl.Update)
        c.DELETE("/:id", ctrl.Delete)
    }

    // c := r.Group("/color")
    // {
    //     ctrl := controllers.CalController{}
    //     c.GET("", ctrl.Index)
    //     c.POST("", ctrl.Create)
    //     c.GET("/:id", ctrl.Show)
    //     c.PUT("/:id", ctrl.Update)
    //     c.DELETE("/:id", ctrl.Delete)
    // }

    return r
}

func setCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET","POST","PUT", "PATCH","DELETE"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}