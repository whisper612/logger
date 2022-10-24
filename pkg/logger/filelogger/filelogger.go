package filelogger

import (
	"os"

	lg "github.com/whisper612/logger/pkg/logger"
)

const (
	log   = "LOG"
	warn  = "WARN"
	err   = "ERROR"
	fatal = "FATAL"
)

type FileLogger struct {
	label            string
	prefixDate       string
	prefixDateFormat string
}

func (l *FileLogger) internalPrintToFile(log string, label string) error {
	l.prefixDate = lg.SetDatePrefix(l.prefixDateFormat)
	l.setLabel(label)

	f, err := os.OpenFile("./output/log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(l.prefixDate + " " + label + " " + log + "\n"); err != nil {
		panic(err)
	}

	return err
}

func (l *FileLogger) setLabel(value string) error {
	l.label = value

	return nil
}

func (l *FileLogger) SetDatePrefixFormat(value string) error {
	l.prefixDateFormat = value

	return nil
}

func (l *FileLogger) Log(msg string) error {
	err := l.internalPrintToFile(msg, log)

	return err
}

func (l *FileLogger) Warn(msg string) error {
	err := l.internalPrintToFile(msg, warn)

	return err
}

func (l *FileLogger) Error(msg error) error {
	err := l.internalPrintToFile(msg.Error(), err)

	return err
}

func (l *FileLogger) Fatal(msg error) error {
	err := l.internalPrintToFile(msg.Error(), fatal)

	if err != nil {
		return err
	}

	panic(msg)
}
