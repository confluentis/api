package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/confluentis/api/entities"
    "gopkg.in/gorp.v1"
    "log"
    "strconv"
)

// Get a collectionof tests
func  GetTests(c *gin.Context) {
    var tests []entities.Test
    db, _ := c.MustGet("db").(*gorp.DbMap)

    _, err := db.Select(&tests, "select * from test")

    if err != nil {
       log.Fatalf("Error getting tests", err)
    }

    c.JSON(200, tests)
}

// Get a single test
func GetTest(c *gin.Context) {
    var test entities.Test
    db, _ := c.MustGet("db").(*gorp.DbMap)

    userId, err := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
    if err != nil {
        c.JSON(404, nil)
        return
    }

    if err := db.SelectOne(&test, "select * from test where id = ?", userId); err != nil {
        log.Print(err)
    }

    c.JSON(200, test)
}