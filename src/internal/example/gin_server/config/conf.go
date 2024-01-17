package config

import (
	"github.com/spf13/viper"
	"github.com/xiaoz194/FlyXGo/src/pkg/utils/logutil"
)

var (
	AppMode  string
	HttpPort string
)

func InitConfig() {
	logutil.LogrusObj.Info("start init...")
	viper.AddConfigPath("src/config")
	viper.SetConfigName("values")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		logutil.LogrusObj.Errorf("Error reading config file, falling back to environment variables: %v", err)
	}
	AppMode = viper.GetString("ginServer.AppMode")
	HttpPort = viper.GetString("ginServer.HttpPort")
	logutil.LogrusObj.Info("init success!")
}
