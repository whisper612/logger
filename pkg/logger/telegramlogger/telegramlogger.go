package telegramlogger

import (
	"fmt"

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

func (l *TelegramLogger) internalPrintToTgMessage(log string, label string) error {
	l.prefixDate = lg.SetDatePrefix(l.prefixDateFormat)
	l.setLabel(label)
	// l.prefixDate+" "+label+" "+log

	_, err := l.Bot.SendMessage(294217967, l.prefixDate+" "+label+" "+log, "", 0, false, false)
	if err != nil {
		fmt.Println(err)
	}

	return nil
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
