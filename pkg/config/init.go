package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var config Config

func init() {
	vi := viper.New()
	vi.AddConfigPath(".")      // 文件所在目录
	vi.SetConfigName("config") // 文件名
	vi.SetConfigType("yaml")   // 文件类型

	if err := vi.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}
	err := vi.Unmarshal(&config)
	if err != nil {
		fmt.Println("[Error] config unmarshal failed. error: ", err)
		os.Exit(-1)
	}
}

func GetConfig() Config {
	return config
}
