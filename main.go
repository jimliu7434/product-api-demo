package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"product-api-demo/common/log"
	"product-api-demo/config"

	Model "product-api-demo/model"
	FakeModel "product-api-demo/model/fake"
	Router "product-api-demo/router"
)

// IsDebugMode : debug mode flag
var IsDebugMode *bool

// ConfigFilePath : 設定檔路徑
var ConfigFilePath *string

// ConfigFileType : 設定檔規格
var ConfigFileType string
var model Model.IModel

func init() {
	IsDebugMode = flag.Bool("debugmode", false, "is in debug mode")
	ConfigFilePath = flag.String("f", "./config.yaml", "config file location")
	flag.Parse()
	log.Initialize(*IsDebugMode)

	// 初始化設定檔
	ConfigFileType = "yaml"
	config.Setup(ConfigFileType, *ConfigFilePath)

	// 刪除過期 log
	log.DelLog()

	// 決定使用哪一種 model
	model = FakeModel.New()
}

func main() {
	// 若需開檔寫 log ，此時生成資料夾及檔案
	log.AddFileHook(config.Server, *IsDebugMode)

	// show config to logger
	if *IsDebugMode == true {
		config.Show(log.AccLog)
	}

	// Set gin to be release mode
	gin.SetMode(gin.ReleaseMode)

	// setup router
	handler := Router.SetupRouter(*IsDebugMode, model)

	// Create server
	s := &http.Server{
		Addr:           config.Server.Port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
