package database

import (
	"database/sql"
	"fmt"

	"github.com/IrDeTen/monitoringbot/blocks"

	_ "github.com/mattn/go-sqlite3"
)

// DataBase ...
var DataBase, err = sql.Open("sqlite3", "MonitoringDB.db")

//GetBlocks ...
func GetBlocks() (b []blocks.Block) {

	Blocks, err := DataBase.Query("select * from FuncBlocks")
	if err != nil {
		panic(err)
	}
	defer Blocks.Close()

	for Blocks.Next() {
		i := blocks.Block{}
		err := Blocks.Scan(&i.BlockName, &i.Delay, &i.InRegulations)
		if err != nil {
			fmt.Println(err)
			continue
		}
		b = append(b, i)
	}
	return b
}

// ChangeRegulationsStatus ...
func ChangeRegulationsStatus(inRegulStatus bool, BlockName string) {
	_, err := DataBase.Exec("update FuncBlocks set InRegulations = $1 where BlockName = $2", inRegulStatus, BlockName)
	if err != nil {
		panic(err)
	}
}

//GetUsers ...
func GetUsers() (u []int64) {
	var ChatIDs int64
	Users, err := DataBase.Query("select * from Users")
	if err != nil {
		panic(err)
	}
	for Users.Next() {
		err := Users.Scan(&ChatIDs)
		if err != nil {
			fmt.Println(err)
			continue
		}
		u = append(u, ChatIDs)
	}
	return u
}
