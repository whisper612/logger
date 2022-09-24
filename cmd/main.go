package main

import (
	cl "github.com/whisper612/logger/pkg/logger/consolelogger"
	"github.com/whisper612/logger/internal/someservice"
)

func main() {
	consoleLogger := &cl.ConsoleLogger{}
	// fileLogger := &cl.FileLogger{}

	service0 := someservice.NewService(consoleLogger)
	// service1 := someservice.NewService(fileLogger)

	service0.Serve()
	// service1.Serve()
}

/*
	@TODO
	----------    ДЗ    ----------
	1) дата в логе
	1*) стектрейс в логе
	2) логгер в файл
	3) логгер куда-нибудь на сервер
*/