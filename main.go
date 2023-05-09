package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/Uuq114/WeiwuCache/core"
)

func LoadConfig() {
	// default values
	viper.SetDefault("configDirPath", "config/")
	// read config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("[ERROR] Read config.yml fail")
	}

	for k, v := range viper.AllSettings() {
		fmt.Printf("[key]%s, [value]%s\n", k, v)
	}
}

func main() {
	LoadConfig()

	cache := core.Cache{}
	cache.Init()

	cache.SetWithDefaultExpiration("caocao", "weiwu")
	result := cache.Get("caocao")
	log.Println("result:", result)
}
