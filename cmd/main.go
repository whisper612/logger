package main

import (
	"github.com/whisper612/logger/internal/someservice"
	cl "github.com/whisper612/logger/pkg/logger/consolelogger"
	fl "github.com/whisper612/logger/pkg/logger/filelogger"
)

func main() {
	consoleLogger := &cl.ConsoleLogger{}
	fileLogger := &fl.FileLogger{}
	// telegramLogger := &cl.TelegramLogger{}

	service0 := someservice.NewService(consoleLogger)
	service1 := someservice.NewService(fileLogger)
	// service2 := someservice.NewService(telegramLogger)

	service0.Serve()
	service1.Serve()
	// service2.Serve()
}

/*
	@TODO
	----------    ДЗ    ----------
	1*) стектрейс в логе
	2) логгер в файл
	3) логгер куда-нибудь на сервер
*/
