package blocks

import (
	"time"
)

//Block ...
type Block struct {
	BlockName     string
	LastTime      time.Time
	Delay         time.Duration
	InRegulations bool
}

//GetBlocks - получение блоков из базы с их последним статусом
