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
		Messages: husk.NewTable(new(Message)),
	}
}

func Shutdown() {
	ctx.Messages.Save()
}
