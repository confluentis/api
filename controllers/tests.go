package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/confluentis/api/entities"
    "gopkg.in/gorp.v1"
    "log"
    "strconv"
    "github.com/confluentis/api/generators"
    "time"
)

// Get a collection of tests
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

// Create a new Test
func PostTest(c *gin.Context) {
    var test entities.Test
    db, _ := c.MustGet("db").(*gorp.DbMap)

    if c.BindJSON(&test) == nil {
        // we should store this.
        test.Token = generators.GenerateToken()
        test.Created = time.Now().UnixNano()

        if err := db.Insert(&test); err != nil {
            log.Printf("Error: %s", err)
            c.JSON(404, nil)
            return
        }

        c.JSON(201, test)
    }
}