package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/confluentis/api/entities"
    "gopkg.in/gorp.v1"
    "log"
)

func  GetTests(c *gin.Context) {
    var tests []entities.Test

    db, _ := c.MustGet("db").(*gorp.DbMap)

    _, err := db.Select(&tests, "select * from test")

    if err != nil {
       log.Fatalf("Error getting tests", err)
    }

    c.JSON(200, tests)
}
