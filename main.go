package main

import (
	"log"
	"time"

	"github.com/m-alfan/bot-jadwal-shalat/config"
	"github.com/m-alfan/bot-jadwal-shalat/handler"
	"github.com/subosito/gotenv"
	tele "gopkg.in/telebot.v3"
)

func init() {
	_ = gotenv.Load()
}

func main() {
	log.Println("Starting services...")

	cfg := config.NewConfig()
	pref := tele.Settings{
		Token:  cfg.TelegramToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	log.Println("Connecting to telegram...")
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	var (
		menu     = &tele.ReplyMarkup{ResizeKeyboard: true}
		selector = &tele.ReplyMarkup{}

		jadwalToday    = menu.Text("Jadwal Shalat Hari Ini")
		jadwalTomorrow = menu.Text("Jadwal Shalat Besok")

		btnPrev = selector.Data("<-", "sebelumnya")
		btnNext = selector.Data("->", "lanjut")
	)

	menu.Reply(menu.Row(jadwalToday, jadwalTomorrow))
	selector.Inline(selector.Row(btnPrev, btnNext))

	log.Println("Register Handler...")
	jadwalHandler := handler.NewJadwalHandler(menu)
	b.Handle("/start", jadwalHandler.Start)
	b.Handle(tele.OnText, jadwalHandler.GetJadwal)

	log.Println("Bot Ready...")
	b.Start()
}
