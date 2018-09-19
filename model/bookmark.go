package model

import (
	"errors"
	"fmt"

	Db "reader_api/database"

	"github.com/jinzhu/gorm"
)

// Bookmark 書籤基本類
type Bookmark struct {
	gorm.Model
	// Book books.Book `gorm:"foreignkey:Book_id"`
	UserID uint `gorm:"unique_index:user_book_uk"`
	BookID uint `gorm:"unique_index:user_book_uk"`
	Chap   uint `gorm:"default:0"`
	Line   uint `gorm:"default:0"`
}

func init() {
	Db.Orm.AutoMigrate(&Bookmark{})
	// db.Model(&BookMark).Related(&book.Book)
}

// 修改
// 刪除

// List 顯示該 userID 書籤列表
func (bookmark Bookmark) List(userID string) (bookmarks []Bookmark, err error) {
	if err = Db.Orm.Where("user_id = ?", userID).Find(&bookmarks).Error; err != nil {
		fmt.Println("[Model Error]")
		fmt.Println(err)
		return
	}

	return
}

// Create 新增書籤
func (bookmark *Bookmark) Create() (err error) {

	if Db.Orm.NewRecord(bookmark) == false {
		err = errors.New("Not a New Record")
		fmt.Println(err)
		return
	}

	if err = Db.Orm.Create(&bookmark).Error; err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	return
}

// Get 撈取指定的資料
func (bookmark *Bookmark) Get() (err error) {
	if err = Db.Orm.Where(&bookmark).First(&bookmark).Error; err != nil {
		fmt.Println(err)
		return
	}
	return
}

// Update 指定目標更新內容
// func (bookmark *Bookmark) Update(userID string, bookID string) (updateResult Bookmark, err error) {

// tar := Db.Orm.Where("user_id = ?", userID).Where("book_id = ?", bookID).First()
// if err = Db.Orm.First(updateResult, "user_id = ?", userID, "book_id = ?", bookID).Error; err != nil {
// fmt.Println(err)
// return
// }

// if Db.Orm.NewRecord(tar) {
// }

// if err =
// }

// func (bookmark *Bookmark) Delete() (err error) {

// }

// func (bookmark *Bookmark) Where() (bookmark Bookmark, err error) {

// }
