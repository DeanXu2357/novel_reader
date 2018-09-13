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

type Bookmark struct {
    gorm.Model
    // Book books.Book `gorm:"foreignkey:Book_id"`
    User_id uint
    Book_id uint
    Chap uint `gorm:"default:0"`
    Line uint `gorm:"default:0"`
}

// config 初始
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

// DB 初始
func init() {
    db_info := "host=" + viper.GetString("DB_HOST")
    db_info += " port=" + viper.GetString("DB_PORT")
    db_info += " user=" + viper.GetString("DB_USER")
    db_info += " dbname=" + viper.GetString("DB_NAME")
    db_info += " password=" + viper.GetString("DB_PASS")

    db, err = gorm.Open("postgres", db_info)

    if err != nil {
        fmt.Println("DATABASE ERROR !!")
        fmt.Println(err)
        fmt.Println("Connecting Data")
        fmt.Println(db_info)
    }
    defer db.Close()

    fmt.Println("[Debug]")
    fmt.Printf("%+v\n", db)
    db.AutoMigrate(&Bookmark{})
    // db.Model(&BookMark).Related(&book.Book)
}

// 新增書籤
func Add(bookmark Bookmark) {
    fmt.Println("debug !!!!")
    fmt.Printf("%+v\n", bookmark)

    fmt.Println("[Debug]")
    fmt.Printf("%+v\n", db)

    if db.NewRecord(bookmark) == false {
        panic("not new record")
    }

    if err := db.Create(&bookmark).Error; err != nil {
        fmt.Printf("%+v\n", err)
        panic(err)
    }
}

// 取得該User的所有書籤
func ListByUser(user_id string) *gorm.DB { //  ([]Bookmark, []error) {
    // var bookmarks []Bookmark

    // err := db.Where("user_id = ?", user_id).Find(&bookmarks).GetErrors

    // return bookmarks, err

    db.DB().Ping()

    return db.Where("user_id = ?", user_id)
}

// 取得指定的書籤（用User&Book查找
func FindByUserAndBook() {
    // code...
}

// 刪除指定書籤
func Delete() {
    // code...
}

// 編輯指定書籤的 chap & line
func Edit() {
    // code...
}

