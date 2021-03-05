package monitor

import (
	"time"

	"github.com/IrDeTen/monitoringbot/bot"
	"github.com/IrDeTen/monitoringbot/database"
)

//TimeMonitoring ...
func TimeMonitoring(b *bot.Bot) {
	//цикл для тестов

	for {

		time.Sleep(10 * time.Second)
		for _, count := range database.GetBlocks() {
			if (count.InRegulations == true) && (time.Now().Sub(count.LastTime) > (count.Delay * time.Second)) {
				database.ChangeRegulationsStatus(!count.InRegulations, count.BlockName)
				b.Message(!count.InRegulations, count.BlockName)
			}
		}
	}
}
