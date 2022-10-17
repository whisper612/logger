package logger

import "github.com/golang-module/carbon/v2"

type Logger interface {
	SetDatePrefixFormat(value uint) error
	Error(err error) error
	Fatal(err error) error
	Warn(log string) error
	Log(log string) error
}

const (
	DateFull = iota
	DayMonthHoursMinute
	DayMonthYear
	DayMonth
	HoursMinuteSeconds
)

func SetDatePrefix(prefixDateFormat uint) string {
	date := carbon.Now().ToDateTimeString()
	switch prefixDateFormat {
	case DateFull:
		return date
	case DayMonthHoursMinute:
		return carbon.Parse(date).Format("d-m H:i")
	case DayMonthYear:
		return carbon.Parse(date).Format("d-m-Y")
	case DayMonth:
		return carbon.Parse(date).Format("d-m")
	case HoursMinuteSeconds:
		return carbon.Parse(date).Format("H:i:s")
	default:
		return ""
	}
}
