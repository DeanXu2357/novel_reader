package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"reader_api/model"
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

			v1.GET("/:user_id/books/", ListBooks)
			v1.PUT("/:user_id/books/:book_id", AddBookMark)
			v1.DELETE("/:user_id/books/:book_id", DeleteBookMark)
			v1.POST("/:user_id/books/:book_id", ReadBook)
		}
	}

	r.Run(":8080")
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

	//   add_item := model.Bookmark{User_id: uint(user_id64), Book_id: uint(book_id64)}
	// add_item.User_id = uint(user_id64)
	// add_item.Book_id = uint(book_id64)
	// if item := bookmark.Add(add_item);item.Error != nil {
	//     fmt.Println("[Error] bookmark.Add")
	//     fmt.Println(item.Error)
	//     c.AbortWithError(200, item.Error)
	// } else {
	//     c.JSON(200, item)
	// }

	// bookmark.Add(add_item)

	c.JSON(200, userID64+bookID64)
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
