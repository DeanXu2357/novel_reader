package main

import (
    "fmt"
    "strings"
    "github.com/spf13/viper"
)

func init() {
    viper.SetConfigType("yaml")
    viper.SetConfigName("config")
    viper.AddConfigPath("./")
    viper.SetEnvPrefix("dev")
    viper.AutomaticEnv()
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

    if err:= viper.ReadInConfig(); err != nil {
        fmt.Println(err)
    }
}

func main() {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
    fmt.Println(viper.Get("SOURCE_BASE_PATH"))
}
