package main

import (
    "github.com/gin-gonic/gin"
    "github.com/confluentis/api/controllers"
    "gopkg.in/gorp.v1"
)

func DbMiddleware(db *gorp.DbMap) gin.HandlerFunc {
    return func (c *gin.Context) {
        c.Set("db", db)
        c.Next()
    }
}

func main() {
    var db = initDb()
    defer db.Db.Close()

    r := gin.Default()
    r.Use(DbMiddleware(db))

    v1 := r.Group("api/v1")
    {
        v1.GET("/ping", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "pong",
            })
        })

        v1.GET("/tests", controllers.GetTests)
        v1.GET("/tests/:id", controllers.GetTest)
        v1.POST("/tests", controllers.PostTest)
    }

    r.Run() // listen and server on 0.0.0.0:8080
}
