package model

// Book 書籍標準類
type Book struct {
	// gorm.Model
	path       string
	name       string
	author     string
	latestChap string
}

func init() {
	// db.AutoMigrate(&Book{})
}
