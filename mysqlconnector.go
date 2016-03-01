package main

import (
    "gopkg.in/gorp.v1"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "log"
    "github.com/confluentis/api/entities"
    "fmt"
    "os"
)

func initDb() *gorp.DbMap {
    mysqlconnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("API_MYSQL_USER"), os.Getenv("API_MYSQL_PASSWORD"), os.Getenv("API_MYSQL_HOST"), os.Getenv("API_MYSQL_PORT"), os.Getenv("API_MYSQL_DATABASE"))

    db, err := sql.Open("mysql", mysqlconnection)

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