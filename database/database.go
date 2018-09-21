package database

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

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
	db_info := "host=" + viper.GetString("DB_HOST")
	db_info += " port=" + viper.GetString("DB_PORT")
	db_info += " user=" + viper.GetString("DB_USER")
	db_info += " dbname=" + viper.GetString("DB_NAME")
	db_info += " password=" + viper.GetString("DB_PASS")
	db_info += " sslmode=disable"

	Orm, err = gorm.Open("postgres", db_info)

	if err != nil {
		fmt.Println("DATABASE ERROR !!")
		fmt.Println(err)
		fmt.Println("Connecting Data")
		fmt.Println(db_info)
	}
	// defer Orm.Close()
}
