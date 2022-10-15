package consolelogger

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/golang-module/carbon/v2"
	lg "github.com/whisper612/logger/pkg/logger"
)

const (
	log   = "LOG"
	warn  = "WARN"
	err   = "ERROR"
	fatal = "FATAL"
)

type ConsoleLogger struct {
	label            string
	prefixDate       string
	prefixDateFormat uint
}

func (l *ConsoleLogger) internalPrint(log string, label string) error {
	l.setDatePrefix()
	l.setLabel(label)

	c := color.New(color.FgWhite)
	switch l.label {
	case warn:
		c = color.New(color.FgYellow)
	case err:
		c = color.New(color.FgRed)
	case fatal:
		c = color.New(color.FgHiMagenta)
	}

	fmt.Printf(l.prefixDate)
	c.Printf(" [%v] ", l.label)
	fmt.Printf("%v\n", log)

	return nil
}

func (l *ConsoleLogger) setLabel(value string) error {
	l.label = value

	return nil
}

func (l *ConsoleLogger) setDatePrefix() error {
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

func (l *ConsoleLogger) SetDatePrefixFormat(value uint) error {
	l.prefixDateFormat = value

	return nil
}

func (l *ConsoleLogger) Log(msg string) error {
	err := l.internalPrint(msg, log)

	return err
}

func (l *ConsoleLogger) Warn(msg string) error {
	err := l.internalPrint(msg, warn)

	return err
}

func (l *ConsoleLogger) Error(msg error) error {
	err := l.internalPrint(msg.Error(), err)

	return err
}

func (l *ConsoleLogger) Fatal(msg error) error {
	err := l.internalPrint(msg.Error(), fatal)

	if err != nil {
		return err
	}

	panic(msg)
}
