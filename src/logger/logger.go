package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func init() {
	Logger = logrus.New()

	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logFile := "/var/log/k8s-api.log"

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Logger.Fatal("Failed to open log file: ", err)
	}

	Logger.SetOutput(io.MultiWriter(file, os.Stdout))

	// TODO: file or Stdout 선택
	// Logger.SetOutput(file)
	// Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.DebugLevel)
}

func Errorf(reqId string, format string, v ...interface{}) {
	loggerWithField := Logger.WithFields(logrus.Fields{})
	loggerWithField.Data["file"] = fileInfo(2)
	if reqId != "" {
		loggerWithField.Data["requestId"] = reqId
	}
	loggerWithField.Errorf(format, v...)
}

func Debugf(reqId string, format string, v ...interface{}) {
	loggerWithField := Logger.WithFields(logrus.Fields{})
	loggerWithField.Data["file"] = fileInfo(2)

	loggerWithField.Data["requestId"] = reqId

	loggerWithField.Debugf(format, v...)
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
