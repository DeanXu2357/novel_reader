package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	// ErrRecordNotFound gorm 的空值錯誤
	ErrRecordNotFound = gorm.ErrRecordNotFound
	// ErrDataExisted 資料已存在
	ErrDataExisted = errors.New("Data existed")
)
