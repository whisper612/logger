package consolelogger

import (
	"fmt"
	"github.com/fatih/color"
)

const (
	log = "LOG"
	warn = "WARN"
	err = "ERROR"
	fatal = "FATAL"
)

type ConsoleLogger struct {
	label string
	prefix string
}

func (l ConsoleLogger) internalPrint(log string, label string) error {
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

	c.Printf("[%v] ", l.label)
	fmt.Printf("%v\n", log)

	return nil
} 

func (l *ConsoleLogger) setLabel(value string) error {
	l.label = value

	return nil
}

func (l *ConsoleLogger) SetPrefix(value string) error {
	l.prefix = value

	return nil
}

func (l ConsoleLogger) Log(msg string) error {
	err := l.internalPrint(msg, log)

	return err
}

func (l ConsoleLogger) Warn(msg string) error {
	err := l.internalPrint(msg, warn)

	return err
}

func (l ConsoleLogger) Error(msg error) error {
	err := l.internalPrint(msg.Error(), err)

	return err
}

func (l ConsoleLogger) Fatal(msg error) error {
	err := l.internalPrint(msg.Error(), fatal)

	if err != nil {
		return err
	}

	panic(msg)
}