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

	if err = Db.Orm.Where("user_id = ?", bookmark.UserID).Where("book_id = ?", bookmark.BookID).First(&bookmark).Error; err != nil {
		fmt.Println(err)
		return
	}
	return
}

// UpdateDetail 指定目標更新內容
func (bookmark *Bookmark) UpdateDetail(chap uint, line uint) (err error) {

	// if bookmark.(type) != Bookmark {
	//     err = fmt.Errorf("Invalid type")
	//     return
	// }
	if bookmark.Chap == chap && bookmark.Line == line {
		return
	}

	if Db.Orm.NewRecord(bookmark) {
		err = ErrRecordNotFound
		return
	}

	if err = Db.Orm.Model(&bookmark).Updates(map[string]interface{}{"chap": chap, "line": line}).Error; err != nil {
		return
	}
	return
}

// Delete 刪除指定目標
func (bookmark *Bookmark) Delete() (err error) {

	if Db.Orm.NewRecord(bookmark) {
		err = ErrRecordNotFound
		return
	}

	if err = Db.Orm.Delete(&bookmark).Error; err != nil {
		return
	}

	return
}

// Book 取得書本資訊
func (bookmark Bookmark) Book() (book Book, err error) {

	if Db.Orm.NewRecord(bookmark) {
		err = ErrRecordNotFound
		return
	}

	if err = Db.Orm.Model(bookmark).Related(&book).Error; err != nil {
		return
	}

	return
}
