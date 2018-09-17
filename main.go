package main

import (
	"fmt"
	"strconv"
	"strings"

	"reader_api/model"
    Log "reader_api/logs"

    "github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("dev")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/test", Test)

	reader := r.Group("/reader")
	{
		v1 := reader.Group("/v1")
		{
			// test route
			// v1.GET("/test", Test)
            v1.Use(LogMiddleware())
			v1.GET("/:user_id/books/", ListBooks)
			v1.PUT("/:user_id/books/:book_id", AddBookMark)
			v1.DELETE("/:user_id/books/:book_id", DeleteBookMark)
			v1.POST("/:user_id/books/:book_id", ReadBook)
		}
	}

	r.Run(":8080")
}

func LogMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        Log.Info.Println(c.Request.URL)
    }
}

// Test 測試路由
func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

// ListBooks 該user_id所有書籤列表
func ListBooks(c *gin.Context) {
	userID := c.Params.ByName("user_id")

	var bookmark model.Bookmark
	bookmarks, err := bookmark.List(userID)
	if err != nil {
		fmt.Println("[Error]")
		fmt.Println(err)
		c.AbortWithError(404, err)
		return
	}

	c.JSON(200, bookmarks)
}

// AddBookMark 添加書籤
func AddBookMark(c *gin.Context) {
	userID := c.Params.ByName("user_id")
	bookID := c.Params.ByName("book_id")
	userID64, _ := strconv.ParseUint(userID, 10, 64)
	bookID64, _ := strconv.ParseUint(bookID, 10, 64)

    // todo 添加檢查 user_book_uk 是否重複

    var bookmark = model.Bookmark{UserID: uint(userID64), BookID: uint(bookID64)}
    if err := bookmark.Create(); err != nil {
        fmt.Println("[Error]")
        fmt.Println(err)
        c.AbortWithError(404, "Sorry, There's something wrong . We'll fix it as soon as possible")
        return
    }

	c.JSON(200, bookmark)
}

// DeleteBookMark 刪除書籤
func DeleteBookMark(c *gin.Context) {
	userID := c.Params.ByName("user_id")
	bookID := c.Params.ByName("book_id")

	c.JSON(200, "user_id = "+userID+", book_id = "+bookID)
}

// ReadBook 從書籤處閱讀
func ReadBook(c *gin.Context) {
	userID := c.Params.ByName("user_id")
	bookID := c.Params.ByName("book_id")

	c.JSON(200, "user_id = "+userID+", book_id = "+bookID)
}
