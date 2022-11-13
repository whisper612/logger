package telegramlogger

import (
	"fmt"
	"os"

	bt "github.com/SakoDroid/telego"
	lg "github.com/whisper612/logger/pkg/logger"
)

const (
	log   = "LOG"
	warn  = "WARN"
	err   = "ERROR"
	fatal = "FATAL"
)

type TelegramLogger struct {
	label            string
	prefixDate       string
	prefixDateFormat string
	Bot              *bt.Bot
}

func (l *TelegramLogger) SendMessage() error {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return err
	}
	logFile, err := os.Open(path + "/output/log.txt")
	if err != nil {
		fmt.Println(err)
		return err
	}

	msg := l.Bot.SendDocument(294217967, 0, "Your's logs, Sir", "")
	_, err = msg.SendByFile(logFile, false, false)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (l *TelegramLogger) internalPrintToTgMessage(log string, label string) error {
	l.prefixDate = lg.SetDatePrefix(l.prefixDateFormat)
	l.setLabel(label)

	textLog := l.prefixDate + " " + label + " " + log + "\n"
	err := lg.PrintToFile(textLog)

	return err
}

func (l *TelegramLogger) setLabel(value string) error {
	l.label = value

	return nil
}

func (l *TelegramLogger) SetDatePrefixFormat(value string) error {
	l.prefixDateFormat = value

	return nil
}

func (l *TelegramLogger) Log(msg string) error {
	err := l.internalPrintToTgMessage(msg, log)

	return err
}

func (l *TelegramLogger) Warn(msg string) error {
	err := l.internalPrintToTgMessage(msg, warn)

	return err
}

func (l *TelegramLogger) Error(msg error) error {
	err := l.internalPrintToTgMessage(msg.Error(), err)

	return err
}

func (l *TelegramLogger) Fatal(msg error) error {
	err := l.internalPrintToTgMessage(msg.Error(), fatal)

	if err != nil {
		return err
	}

	panic(msg)
}
