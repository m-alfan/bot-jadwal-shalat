package handler

import (
	"log"

	tele "gopkg.in/telebot.v3"
)

type jadwal struct {
	teleMenu *tele.ReplyMarkup
}

type JadwalHandler interface {
	Start(c tele.Context) error
	GetJadwal(c tele.Context) error
}

func NewJadwalHandler(teleMenu *tele.ReplyMarkup) *jadwal {
	return &jadwal{teleMenu}
}

func (j *jadwal) Start(c tele.Context) error {
	log.Println("Incoming initial message")
	return c.Send("Assalamu 'alaikum...", j.teleMenu)
}

func (j *jadwal) GetJadwal(c tele.Context) error {
	log.Println("Incoming text message")
	return c.Send(c.Text())
}
