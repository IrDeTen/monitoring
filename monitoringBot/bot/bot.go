package bot

import (
	"database/sql"
	"fmt"

	"github.com/IrDeTen/monitoringbot/database"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	_ "github.com/mattn/go-sqlite3"
)

// Bot ...
type Bot struct {
	BotApi   *tgbotapi.BotAPI
	BotToken string
}

//StartBot ...
func (b *Bot) StartBot(DB *sql.DB) {

	Config := DB.QueryRow("select * from Config")

	err := Config.Scan(&b.BotToken)
	if err != nil {
		panic(err)
	}
	fmt.Println(b.BotToken)

	BotAPI, err := tgbotapi.NewBotAPI(b.BotToken)
	if err != nil {
		panic(err)
	}
	(*b).BotApi = BotAPI
	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 60
	updates, _ := b.BotApi.GetUpdatesChan(upd)
	for {
		for update := range updates {
			ChatID := update.Message.Chat.ID
			msg := tgbotapi.NewMessage(ChatID, "It's working")
			b.BotApi.Send(msg)
		}
	}
}

//Message ...
func (b *Bot) Message(inRegulStatus bool, BlockName string) {
	for _, count := range database.GetUsers() {
		if inRegulStatus == false {
			Text := "Блок " + BlockName + "вышел из регламента"
			msg := tgbotapi.NewMessage(count, Text)
			b.BotApi.Send(msg)
		}
	}
}
