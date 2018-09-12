package book

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Book struct {
    gorm.Model
    path string
    name string
    author string
    latest_chap string
}

func init() {
    db, err = gorm.Open("postgres", "host= port= user= dbname= password=")

    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

    db.AutoMigrate(&Book{})
}


