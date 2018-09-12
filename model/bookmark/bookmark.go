package bookmark

import (
    "fmt"
    "strings"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/spf13/viper"
    
    // "reader_api/model/book"
)

var db *gorm.DB
var err error

type BookMark struct {
    gorm.Model
    // Book books.Book `gorm:"foreignkey:Book_id"`
    User_id uint
    Book_id uint
    Chap uint
    Line uint
}

func init() {
    viper.SetConfigType("yaml")
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    viper.SetEnvPrefix("dev")
    viper.AutomaticEnv()
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

    if err:= viper.ReadInConfig(); err != nil {
        fmt.Println(err)
    }
}

func init() {
    db_info := "host=" + viper.GetString("DB_HOST") 
    db_info += " port=" + viper.GetString("DB_PORT") 
    db_info += " user=" + viper.GetString("DB_USER") 
    db_info += " dbname=" + viper.GetString("DB_NAME") 
    db_info += " password=" + viper.GetString("DB_PASS")

    db, err = gorm.Open("postgres", db_info)

    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

    db.AutoMigrate(&BookMark{})
    // db.Model(&BookMark).Related(&book.Book)
}

func Add() {
    // code...
}

func ListByUser(user_id) {
    // code...
}

