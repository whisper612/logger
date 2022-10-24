package logger

import "github.com/golang-module/carbon/v2"

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
