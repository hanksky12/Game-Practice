package log

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/*
ex:

	log.WithFields(log.Fields{
	  "event": event,
	  "topic": topic,
	  "key": key,
	}).Fatal("Failed to send event")
*/

func Init(filePath string, level string) {
	env := "development" //development //large_test
	setLevel(env, level)
	setFormatter(env)
	//setOutput(filePath)
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
		LocalTime:  true,
	}
	// 設置文件權限為 0644
	go func() {
		time.Sleep(1 * time.Second)
		_, err := os.Stat(filePath)
		if err == nil {
			err = os.Chmod(filePath, 0644)
			if err != nil {
				log.Fatalf("Failed to change log file permissions: %v", err)
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
			gid := getGoroutineID()
			return fmt.Sprintf("[%s] %s", gid, frame.Function), fileName
		}})
}

func getGoroutineID() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	if id, err := strconv.Atoi(idField); err == nil {
		return fmt.Sprintf("Goroutine-%d", id)
	}
	return "Goroutine-Unknown"
}

func setLevel(env, level string) {
	if level != "" {
		switch level {
		case "debug":
			log.SetLevel(log.DebugLevel)
		case "info":
			log.SetLevel(log.InfoLevel)
		case "warn":
			log.SetLevel(log.WarnLevel)
		default:
			log.SetLevel(log.InfoLevel) // 默認使用 info 級別
		}
		return
	}
	switch env {
	case "development":
		log.SetLevel(log.DebugLevel)
	case "test", "production":
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.InfoLevel) // 默認使用 info 級別
	}
}
