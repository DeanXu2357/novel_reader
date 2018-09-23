package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	Log "reader_api/logs"
	"reader_api/model"

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

// LogMiddleware 紀錄資料中介層
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		Log.Info.Println("=============================")
		Log.Info.Println("=== Request Input         ===")
		Log.Info.Println(c.Request.URL)

		Log.Info.Println("=== Request Header        ===")
		for k, v := range c.Request.Header {
			Log.Info.Println(k, v)
		}

		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

		Log.Info.Println("=== Request Body          ===")
		body, _ := ioutil.ReadAll(rdr1)
		Log.Info.Println(string(body))
		c.Request.Body = rdr2
		// Log.Info.Println("=============================")

		rl := &responseLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = rl

		c.Next()

		Log.Info.Println("=== Response Body         ===")
		Log.Info.Println(rl.body.String())
		Log.Info.Println("=============================")
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
		Log.Error.Println(err)
		c.AbortWithError(404, err)
		return
	}

	c.JSON(200, bookmarks)
}

// AddBookMark 添加書籤
func AddBookMark(c *gin.Context) {
	userID := c.Params.ByName("user_id")
	bookID := c.Params.ByName("book_id")
	chap := c.PostForm("chap")
	line := c.PostForm("line")
	userID64, _ := strconv.ParseUint(userID, 10, 64)
	bookID64, _ := strconv.ParseUint(bookID, 10, 64)
	chap64, _ := strconv.ParseUint(chap, 10, 64)
	line64, _ := strconv.ParseUint(line, 10, 64)

	bookmark := model.Bookmark{UserID: uint(userID64), BookID: uint(bookID64), Chap: uint(chap64), Line: uint(line64)}
	err := bookmark.Get()

	switch err {
	case nil:
		if errU := bookmark.UpdateDetail(uint(chap64), uint(line64)); errU != nil {
			fmt.Println("[Error]")
			fmt.Println(errU)
			Log.Error.Println(errU)
			c.AbortWithError(404, errU)
			return
		}
		c.JSON(200, bookmark)
		return
	case model.ErrRecordNotFound:
		if errC := bookmark.Create(); errC != nil {
			fmt.Println("[Error]")
			fmt.Println(errC)
			Log.Error.Print(errC)
			c.AbortWithError(404, errC)
			return
		}
		c.JSON(200, bookmark)
		return
	default:
		fmt.Println("[Error]")
		fmt.Println(err)
		Log.Error.Println(err)
		c.AbortWithError(404, err)
		return
	}
}

// DeleteBookMark 刪除書籤
func DeleteBookMark(c *gin.Context) {
	userID := c.Params.ByName("user_id")
	bookID := c.Params.ByName("book_id")
	userID64, _ := strconv.ParseUint(userID, 10, 64)
	bookID64, _ := strconv.ParseUint(bookID, 10, 64)

	bookmark := model.Bookmark{UserID: uint(userID64), BookID: uint(bookID64)}
	err := bookmark.Get()
	switch err {
	case nil:
		// TODO DELETE
		if errD := bookmark.Delete(); errD != nil {
			fmt.Println("[Error]")
			fmt.Println(errD)
			Log.Error.Println(errD)
			c.AbortWithError(404, errD)
			return
		}
		c.JSON(200, bookmark)
		return
	case model.ErrRecordNotFound:
		c.JSON(200, "No Data to delete")
		return
	default:
		fmt.Println("[Error]")
		fmt.Println(err)
		Log.Error.Println(err)
		c.AbortWithError(404, err)
		return
	}
}

// ReadBook 從書籤處閱讀
func ReadBook(c *gin.Context) {
	userID := c.Params.ByName("user_id")
	bookID := c.Params.ByName("book_id")
	// chap := c.PostForm("chap")
	// line := c.DefaultPostForm("line", "113")
	userID64, _ := strconv.ParseUint(userID, 10, 64)
	bookID64, _ := strconv.ParseUint(bookID, 10, 64)

	bookmark := model.Bookmark{UserID: uint(userID64), BookID: uint(bookID64)}
	bookmark.Get()

	c.JSON(200, bookmark)
	return
}
