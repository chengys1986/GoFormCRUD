package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() error {
	fmt.Println("InitConfig")
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func ReadConfig(inputStr string) string {
	fmt.Printf("ReadConfig: %s\n", inputStr)
	result := viper.GetString(inputStr)
	fmt.Printf("ReadConfi(Result): %s\n", result)
	return viper.GetString(inputStr)
}
