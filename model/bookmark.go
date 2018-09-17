package model

import (
	"fmt"
    "errors"

	Db "reader_api/database"

	"github.com/jinzhu/gorm"
	// "reader_api/model/book"
)

// Bookmark 書籤基本類
type Bookmark struct {
	gorm.Model
	// Book books.Book `gorm:"foreignkey:Book_id"`
	UserID uint
	BookID uint
	Chap   uint `gorm:"default:0"`
	Line   uint `gorm:"default:0"`
}

func init() {
	Db.Orm.AutoMigrate(&Bookmark{})
	// db.Model(&BookMark).Related(&book.Book)
}

// var bookmarks []Bookmark
// var err error

// 新增
// 修改
// 刪除

// List 顯示全部
func (bookmark Bookmark) List(userID string) (bookmarks []Bookmark, err error) {
	if err = Db.Orm.Where("user_id = ?", userID).Find(&bookmarks).Error; err != nil {
		fmt.Println("[Model Error]")
		// fmt.Printf("%+v\n", err)
		fmt.Println(err)
		return
	}

	return
}

// Create 新增書籤
func (bookmark *Bookmark) Create() (err error) {

    if Db.Orm.NewRecord(bookmark) == false {
        err = errors.New("Not a New Record !!")
        fmt.Println(err)
        return
    }

	if err = Db.Orm.Create(&bookmark).Error; err != nil {
	    fmt.Printf("%+v\n", err)
        return
	}

    return
}

// 取得該User的所有書籤
// func ListByUser(user_id string) *gorm.DB { //  ([]Bookmark, []error) {
// var bookmarks []Bookmark

// err := db.Where("user_id = ?", user_id).Find(&bookmarks).GetErrors

// return bookmarks, err

//     db.DB().Ping()

//     return db.Where("user_id = ?", user_id)
// }

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
