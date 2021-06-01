package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var vp *viper.Viper

// Server : config.Server
var Server *ServerConfig

// IConfig : config type 規範
type IConfig interface {
	Show()
	Read()
}

func init() {
	Server = &ServerConfig{}
}

// Setup 初始化設定
func Setup(configFileType string, configFilePath string) {
	// 初始化 viper
	vp = viper.New()
	vp.SetConfigType(configFileType)
	vp.SetConfigFile(configFilePath)

	if err := vp.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("讀取設定檔失敗，請檢查 config.yaml 檔案是否存在: %v", err))
	}

	// 初始化設定
	Server.Read(vp)
}

func Show(log *logrus.Entry) {
	log.Info("========== Config ==========")
	Server.Show(log)
	log.Info("========== Config ==========")
}
