package main

import (
	"log"
	"encoding/json"
	"math/rand"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"flag"
)

type templates struct {
	Curses `json:"curses"`
}
type Curses []string

var data []byte

func init() {
	fileData, err := ioutil.ReadFile("templates.json")
	if err != nil {
		log.Fatalln(err)
	}
	data = fileData
}

func main() {
	var token string
	flag.StringVar(&token, "token", "empty", "telegram bot token")
	flag.Parse()
	if token == "empty" {
		panic(getRandomCurse())
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			var messageText string
			messageText = getRandomCurse()
			message := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
			bot.Send(message)
			continue
		}

		if update.InlineQuery == nil {
			continue
		}

		curse := getRandomCurse()
		var articles []interface{}
		article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Выебать мамку", curse)
		article.Description = "И не только мамку"
		articles = append(articles, article)

		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			IsPersonal:    true,
			CacheTime:     0,
			Results:       articles,
		}

		if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
			log.Println(err)
		}
	}
}

func getRandomCurse() string {
	var templates templates
	if err := json.Unmarshal(data, &templates); err != nil {
		panic(err)
	}
	cursesSize := len(templates.Curses)
	curse := templates.Curses[rand.Intn(cursesSize)]
	return curse
}
