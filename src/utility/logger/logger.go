package logger

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func Init(logFilePath string) error {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return errors.Wrap(err, "open log file")
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

func Infoln(message ...any) {
	infoLogger.Println(message)
}

func Errorln(message ...any) {
	errorLogger.Println(message)
}
