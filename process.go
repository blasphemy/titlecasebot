package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	"log"
	"strings"
)

func processInlineQuery(i *tgbotapi.InlineQuery, b *tgbotapi.BotAPI) {
	if i.Query == "" {
		log.Println("Ignoring empty query")
		return
	}
	id := uuid.New().String()
	queryRes := strings.Title(i.Query)
	log.Printf("New Query: %s", i.Query)
	log.Printf("Query Result: %s", queryRes)
	tString := fmt.Sprintf("Title Case: \"%s\"", queryRes)
	wf := tgbotapi.NewInlineQueryResultArticleMarkdown(id, tString, queryRes)
	icfg := tgbotapi.InlineConfig{
		InlineQueryID: i.ID,
		Results:       []interface{}{wf},
	}
	_, err := b.AnswerInlineQuery(icfg)
	if err != nil {
		log.Println("Error sending result")
		log.Println(err.Error())
	}
}
