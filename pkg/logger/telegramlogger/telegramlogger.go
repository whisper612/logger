package telegramlogger

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
}

func (l *TelegramLogger) internalPrintToFile(log string, label string) error {
	l.prefixDate = lg.SetDatePrefix(l.prefixDateFormat)
	l.setLabel(label)
	// l.prefixDate + " " + label + " " + log

	bot, err := tg.NewBotAPI("828961048:AAFyRStUfD1J9ryLvWYRUb1r1MaTMcNkmjA")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

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
	err := l.internalPrintToFile(msg, log)

	return err
}

func (l *TelegramLogger) Warn(msg string) error {
	err := l.internalPrintToFile(msg, warn)

	return err
}

func (l *TelegramLogger) Error(msg error) error {
	err := l.internalPrintToFile(msg.Error(), err)

	return err
}

func (l *TelegramLogger) Fatal(msg error) error {
	err := l.internalPrintToFile(msg.Error(), fatal)

	if err != nil {
		return err
	}

	panic(msg)
}
