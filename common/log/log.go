package log

import (
	"io"
	"os"
	"product-api-demo/config"
	"time"

	"github.com/sirupsen/logrus"
)

var ErrLogger, AccLogger *logrus.Logger
var ErrLog, AccLog *logrus.Entry
var (
	path    = "_log"
	keepDay = 60
)

func Initialize(isDebugMode bool) {
	AccLogger = logrus.New()
	AccLog = AccLogger.WithFields(logrus.Fields{
		"apitype": "product-api-demo-access",
	})
	ErrLogger = logrus.New()
	ErrLog = ErrLogger.WithFields(logrus.Fields{
		"apitype": "product-api-demo-error",
	})
	if isDebugMode == false {
		AccLogger.SetLevel(logrus.InfoLevel)
	} else {
		AccLogger.SetLevel(logrus.DebugLevel)
	}
	ErrLogger.SetLevel(logrus.ErrorLevel)
	AccLogger.SetOutput(os.Stdout)
	ErrLogger.SetOutput(os.Stdout)
}

func AddFileHook(config *config.ServerConfig, isDebugMode bool) {
	thisDate := time.Now().Format("20060102")
	createFileHook(config, AccLogger, path+"/access_"+thisDate+".log", isDebugMode)
	createFileHook(config, ErrLogger, path+"/error_"+thisDate+".log", isDebugMode)
}

// createFileHook : initialize log file hook
func createFileHook(config *config.ServerConfig, logger *logrus.Logger, outpath string, isDebugMode bool) {
	if config.FileEnabled == true {
		_, err := os.Stat(outpath)
		os.Mkdir(path, os.ModePerm)
		if os.IsNotExist(err) {
			// 檔案不存在，建立新檔
			os.Create(outpath)
		}

		file, err := os.OpenFile(outpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			mw := io.MultiWriter(os.Stdout, file)
			logger.Out = mw
			// logger.Out = file
		} else {
			logger.Error("open " + outpath + " log 檔失敗")
		}
	}
	return
}
