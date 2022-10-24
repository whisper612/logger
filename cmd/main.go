package main

import (
	"github.com/whisper612/logger/internal/someservice"
	cl "github.com/whisper612/logger/pkg/logger/consolelogger"
	fl "github.com/whisper612/logger/pkg/logger/filelogger"
	tg "github.com/whisper612/logger/pkg/logger/telegramlogger"
)

func main() {
	consoleLogger := &cl.ConsoleLogger{}
	fileLogger := &fl.FileLogger{}
	telegramLogger := &tg.TelegramLogger{}

	service0 := someservice.NewService(consoleLogger)
	service1 := someservice.NewService(fileLogger)
	service2 := someservice.NewService(telegramLogger)

	service0.Serve()
	service1.Serve()
	service2.Serve()
}
