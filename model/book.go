package model

import (
	Db "reader_api/database"

	"github.com/jinzhu/gorm"
)

// Book 書籍資訊
type Book struct {
	gorm.Model
	Path       string
	Name       string
	Author     string
	LatestChap string
}

func init() {
	Db.Orm.AutoMigrate(&Book{})
}

// Create 新增書籍
func (book *Book) Create() (err error) {

	if Db.Orm.NewRecord(book) == false {
		err = ErrDataExisted
		return
	}

	if err = Db.Orm.Create(&book).Error; err != nil {
		return
	}

	return
}

// Delete 刪除書籍
func (book *Book) Delete() (err error) {

	if Db.Orm.NewRecord(book) {
		err = ErrRecordNotFound
		return
	}

	if err = Db.Orm.Delete(&book).Error; err != nil {
		return
	}

	return
}

// Get 用傳入的資訊，查找書籍
func (book *Book) Get() (err error) {

	query := book
	if err = Db.Orm.Where(query).First(&book).Error; err != nil {
		return
	}

	return
}

// Update 修改書籍資料
func (book *Book) Update(updateBook Book) (err error) {

	if Db.Orm.NewRecord(book) {
		err = ErrRecordNotFound
		return
	}

	if err = Db.Orm.Model(&book).Updates(updateBook).Error; err != nil {
		return
	}

	return
}
