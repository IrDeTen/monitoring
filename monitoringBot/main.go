package main

import (
	"github.com/IrDeTen/monitoringbot/bot"
	"github.com/IrDeTen/monitoringbot/database"
	"github.com/IrDeTen/monitoringbot/monitor"
)

var NewBot *bot.Bot

func main() {
	NewBot = &bot.Bot{}
	go NewBot.StartBot(database.DataBase)
	monitor.TimeMonitoring(NewBot)
	defer database.DataBase.Close()

}
