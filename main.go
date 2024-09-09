package main

import (
	"bufio"
	"fmt"
	"gopkg.in/telebot.v3"
	"os"
	"time"
)

func main() {
	err := ReadConfig()
	if err != nil {
		panic("read config err" + err.Error())
	}
	fmt.Println("input：")
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  input,
		Poller: &telebot.LongPoller{Timeout: 6 * time.Second},
	})
	if err != nil {
		panic(err)
	}

	bot.Handle("/start", func(c telebot.Context) error {
		groupBtn := telebot.InlineButton{
			Unique: "group",
			Text:   config.TextGroup,
			URL:    config.CaptainGroup,
		}
		xBtn := telebot.InlineButton{
			Unique: "x",
			Text:   config.TextX,
			URL:    config.CaptainX,
		}
		aboutBtn := telebot.InlineButton{
			Unique: "about",
			Text:   config.TextAbout,
		}

		inlineKeys := [][]telebot.InlineButton{
			{groupBtn},
			{xBtn},
			{aboutBtn},
		}

		return c.Send(config.TextWebApp+"\n"+config.CaptainWebApp, &telebot.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		})
	})

	bot.Handle(&telebot.InlineButton{Unique: "about"}, func(c telebot.Context) error {
		return c.Send(config.CaptainAbout)
	})
	fmt.Println("start ...")
	bot.Start()
}
