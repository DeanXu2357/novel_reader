package main

import (
    "fmt"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"

    // "reader_api/model/bookmarks"
)

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

func main () {
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

func Test(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "hello world",
    })
}

func ListBooks(c *gin.Context) {
    var bookmarks []bookmarks.Bookmarks
    user_id := c.Params.ByName("user_id")


    c.JSON(200, "user_id = " + user_id)

    // list := bookmarks.ListByUser(user_id)

    // 使用where in ， 查找list裡面的book_id的書籍資料

    // c.JSON(200, list)
}

func AddBookMark(c *gin.Context) {
    user_id := c.Params.ByName("user_id")
    book_id := c.Params.ByName("book_id")

    c.JSON(200, "user_id = " + user_id + ", book_id = " + book_id)
}

func DeleteBookMark(c *gin.Context) {
    user_id := c.Params.ByName("user_id")
    book_id := c.Params.ByName("book_id")

    c.JSON(200, "user_id = " + user_id + ", book_id = " + book_id)
}

func ReadBook(c *gin.Context) {
    user_id := c.Params.ByName("user_id")
    book_id := c.Params.ByName("book_id")

    c.JSON(200, "user_id = " + user_id + ", book_id = " + book_id)
}

