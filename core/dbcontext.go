package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Messages husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Messages: husk.NewTable(Message{}, serials.GobSerial{}),
	}
}

func Shutdown() {
	ctx.Messages.Save()
}
