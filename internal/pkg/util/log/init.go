package log

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"runtime"
)

/*
ex:

	log.WithFields(log.Fields{
	  "event": event,
	  "topic": topic,
	  "key": key,
	}).Fatal("Failed to send event")
*/

func Init(filePath string) {
	env := os.Getenv("GO_ENV")
	setLevel(env)
	setFormatter(env)
	setOutput(filePath)
	log.SetReportCaller(true)
	log.Info("Log Init Success")
}

func setOutput(filePath string) {
	writer1 := os.Stdout
	writer2 := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    10, // 每個日誌文件最大 10 MB
		MaxBackups: 3,  // 保留最近 3 個日誌文件
		MaxAge:     7,  // 保留 7 天
		Compress:   true,
	}
	// 設置文件權限為 0644
	go func() {
		for {
			_, err := os.Stat(filePath)
			if err == nil {
				err = os.Chmod(filePath, 0644)
				if err != nil {
					log.Fatalf("Failed to change log file permissions: %v", err)
				}
			}
		}
	}()
	log.SetOutput(io.MultiWriter(writer1, writer2))

}

func setFormatter(env string) {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) { //自定义Caller的返回
			fileName := path.Base(frame.File)
			return frame.Function, fileName
		}})
}

func setLevel(env string) {
	switch env {
	case "development":
		log.SetLevel(log.DebugLevel)
	case "test", "production":
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.InfoLevel) // 默認使用 info 級別
	}
}
