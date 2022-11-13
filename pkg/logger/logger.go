package logger

import (
	"errors"
	"os"

	"github.com/golang-module/carbon/v2"
)

type Logger interface {
	SetDatePrefixFormat(value string) error
	Error(err error) error
	Fatal(err error) error
	Warn(log string) error
	Log(log string) error
}

func SetDatePrefix(prefixDateFormat string) string {
	date := carbon.Now().ToDateTimeString()
	if prefixDateFormat != "" {
		return carbon.Parse(date).Format(prefixDateFormat)
	} else {
		return date
	}
}

func PrintToFile(logText string) error {
	if _, err := os.Stat("./output/log.txt"); err == nil {
		file, err := os.OpenFile("./output/log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		defer file.Close()
		if _, err = file.WriteString(logText); err != nil {
			return err
		}
	} else if errors.Is(err, os.ErrNotExist) {
		// вот тут костыль
		// TODO сделать проверку директории
		// TODO в дебаге open w:\Code\Go\github.com\whisper612\logger\cmd/output/log.txt: The system cannot find the path specified.
		file, err := os.Create("./output/log.txt")
		defer file.Close()
		if _, err = file.WriteString(logText); err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}
