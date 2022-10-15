package filelogger

import (
	"os"
	"strings"

	"github.com/golang-module/carbon/v2"
	lg "github.com/whisper612/logger/pkg/logger"
)

const (
	log   = "LOG"
	warn  = "WARN"
	err   = "ERROR"
	fatal = "FATAL"
)

type FileLogger struct {
	label      string
	prefixDate string

	prefixDateFormat uint
}

func (l *FileLogger) internalPrintToFile(log string, label string) error {
	l.setDatePrefix()
	l.setLabel(label)

	f, err := os.OpenFile("../output/log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	defer f.Close()

	if _, err = f.WriteString(l.prefixDate + " " + label + "\n"); err != nil {
		panic(err)
	}

	return err
}

func (l *FileLogger) setLabel(value string) error {
	l.label = value

	return nil
}

func (l *FileLogger) setDatePrefix() error {
	switch l.prefixDateFormat {
	case lg.DateFull:
		l.prefixDate = carbon.Now().ToDateTimeString()
	case lg.DayMonthHoursMinute:
		date := strings.Split(carbon.Now().ToDateTimeString(), "-")
		day := strings.Split(date[2], " ")[0]
		time := strings.Split(carbon.Now().ToDateTimeString(), ":")
		hour := strings.Split(time[0], " ")[1]
		l.prefixDate = date[1] + "-" + day + " " + hour + ":" + time[1]
	case lg.DayMonthYear:
		date := strings.Split(carbon.Now().ToDateTimeString(), "-")
		day := strings.Split(date[2], " ")[0]
		l.prefixDate = day + "-" + date[1] + "-" + date[0]
	case lg.DayMonth:
		date := strings.Split(carbon.Now().ToDateTimeString(), "-")
		day := strings.Split(date[2], " ")[0]
		l.prefixDate = date[1] + "-" + day
	case lg.HoursMinuteSeconds:
		l.prefixDate = carbon.Now().ToTimeString()

	}

	return nil
}

func (l *FileLogger) SetDatePrefixFormat(value uint) error {
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
