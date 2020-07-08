package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Messages husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Messages: husk.NewTable(Message{}),
	}
}

func Shutdown() {
	ctx.Messages.Save()
}
