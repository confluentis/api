package main

import (
    "github.com/gin-gonic/gin"
)

var db = initDb()

func main() {
    r := gin.Default()

    v1 := r.Group("api/v1")
    {
        v1.GET("/ping", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "pong",
            })
        })
    }

    r.Run() // listen and server on 0.0.0.0:8080
}
