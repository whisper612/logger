package main

import (
	"os"

	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
	"github.com/whisper612/logger/internal/someservice"
	cl "github.com/whisper612/logger/pkg/logger/consolelogger"
	fl "github.com/whisper612/logger/pkg/logger/filelogger"
	tg "github.com/whisper612/logger/pkg/logger/telegramlogger"
)

func main() {
	bot, err := bt.NewBot(cfg.Default(os.Getenv("TELEGRAM_APITOKEN")))
	if err != nil {
		panic(err)
	}
	err = bot.Run()
	if err != nil {
		panic(err)
	}

	consoleLogger := &cl.ConsoleLogger{}
	fileLogger := &fl.FileLogger{}
	telegramLogger := &tg.TelegramLogger{
		Bot: bot,
	}

	service0 := someservice.NewService(consoleLogger)
	service1 := someservice.NewService(fileLogger)
	service2 := someservice.NewService(telegramLogger)

	service0.Serve()
	service1.Serve()
	service2.Serve()
}
