package main

import (
	"log"
	"time"

	"github.com/m-alfan/bot-jadwal-shalat/config"
	"github.com/subosito/gotenv"
	tele "gopkg.in/telebot.v3"
)

func init() {
	_ = gotenv.Load()
}

func main() {
	cfg := config.NewConfig()
	pref := tele.Settings{
		Token:  cfg.TelegramToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
