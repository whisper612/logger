package logger

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