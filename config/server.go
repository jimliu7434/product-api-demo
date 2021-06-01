package config

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 資料庫設定
type ServerConfig struct {
	ActiveMsg   string
	Port        string
	FileEnabled bool
}

func (o *ServerConfig) Show(log *logrus.Entry) {
	root := reflect.ValueOf(o).Elem()
	rootTypeName := root.Type().Name()
	for i := 0; i < root.NumField(); i++ {
		propTypeName := root.Type().Field(i).Name
		propValue := root.Field(i).Interface()
		st := fmt.Sprintf("%s.%s=%v", rootTypeName, propTypeName, propValue)
		log.Info(st)
	}
}

func (o *ServerConfig) Read(viper *viper.Viper) {
	o.ActiveMsg = viper.GetString("server.activeMsg")
	o.Port = viper.GetString("server.port")
	o.FileEnabled = viper.GetBool("server.fileEnabled")
}
