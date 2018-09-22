package database

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	// postgres db driver import
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

// Orm gorm基本Db
var Orm *gorm.DB
var err error

// config 初始
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

// DB 初始
func init() {
	dbInfo := "host=" + viper.GetString("DB_HOST")
	dbInfo += " port=" + viper.GetString("DB_PORT")
	dbInfo += " user=" + viper.GetString("DB_USER")
	dbInfo += " dbname=" + viper.GetString("DB_NAME")
	dbInfo += " password=" + viper.GetString("DB_PASS")
	dbInfo += " sslmode=disable"

	Orm, err = gorm.Open("postgres", dbInfo)

	if err != nil {
		fmt.Println("DATABASE ERROR !!")
		fmt.Println(err)
		fmt.Println("Connecting Data")
		fmt.Println(dbInfo)
	}
}
