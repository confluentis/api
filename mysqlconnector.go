package main

import (
    "gopkg.in/gorp.v1"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "log"
    "github.com/confluentis/api/entities"
)

func initDb() *gorp.DbMap {
    db, err := sql.Open("mysql", "root:123@tcp(192.168.33.99:3306)/confluentis")

    if err != nil {
        log.Fatal("Failed connection", err)
    }

    dbmap := &gorp.DbMap{
        Db: db,
        Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"},
    }

    // Add tables
    dbmap.AddTableWithName(entities.Test{}, "test").SetKeys(true, "Id")
    dbmap.AddTableWithName(entities.Task{}, "task").SetKeys(true, "Id")

    if err = dbmap.CreateTablesIfNotExists(); err != nil {
        log.Fatal("Failed creating tables ", err)
    }

    return dbmap
}