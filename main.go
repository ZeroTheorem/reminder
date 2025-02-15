package main

import (
	tele "gopkg.in/telebot.v4"
	"log"
	"math/rand"
	"time"
)

const (
	startMessage = `
	Я создан что бы напоминать тебе мазать ручки кремом! Буду напоминать каждый день в 23:59.
Только попробуй меня ослушаться, тебе пизда!❤️
`
	msg1 = `
	Ахуела сейчас??? А ну бегом мазать ручки, гадина!
`
	msg2 = `
	Cовсем страх потеряла, намазала ручки быстро!
`
	msg3 = `
	Не то нажала, дурында, хорошо подумай, намаж руки и исправь	это недоразумение
`
	msg4 = `
	Патороч, ответ не принят, бегом руки мазать!
`
	msg5 = `
	Cмелая я посмотрю, пиздуй за кремом!
`
	msg6 = `Та мне похуй, что ты там нажала, давай по новой, все хуйня!`
)

func main() {
	pref := tele.Settings{
		Token:  "7661864692:AAETirYrxakprt86ITrRcTQQruoV3ssdnYI",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	loc := time.FixedZone("UTC+4", 4*60*60) // 4 часа * 60 минут * 60 секунд

	s := [4]string{"❤️", "💜", "💙", "💚"}
	ph := [6]string{msg1, msg2, msg3, msg4, msg5, msg6}
	m := &tele.ReplyMarkup{}
	y := m.Data("Да", "1")
	n := m.Data("Нет (даже не думай)", "2")
	m.Inline(m.Row(y), m.Row(n))

	b.Handle("/start", func(c tele.Context) error {
		c.Send(startMessage)
		w := while(loc)
		time.Sleep(time.Duration(w))
		return c.Send("Эй, время мазать ручки! Намазала?", m)
	})
	b.Handle(&y, func(c tele.Context) error {
		c.Send("Твои ручки скажут тебе спасибо, я люблю тебя")
		c.Send(s[rand.Intn(len(s)-1)])
		w := while(loc)
		time.Sleep(time.Duration(w))
		return c.Send("Эй, время мазать ручки! Намазала?", m)
	})

	b.Handle(&n, func(c tele.Context) error {
		return c.Send(ph[rand.Intn(len(ph)-1)], m)
	})
	b.Start()
}

func while(loc *time.Location) int64 {
	now := time.Now().In(loc)
	targetTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, loc)
	diff := targetTime.Sub(now).Nanoseconds()
	return diff

}
