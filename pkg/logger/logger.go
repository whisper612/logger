package logger

type Logger interface {
	SetPrefix(value string) error
	Error(err error) error
	Fatal(err error) error
	Warn(log string) error
	Log(log string) error
}