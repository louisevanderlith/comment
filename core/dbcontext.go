package core

import (
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
)

type CommentContext interface {
}

type context struct {
	Messages husk.Table
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

func GetMessageByKey(key hsk.Key) (hsk.Record, error) {
	return ctx.Messages.FindByKey(key)
}

func GetNodeMessages(itemKey hsk.Key, commentType commenttype.Enum) (records.Page, error) {
	return ctx.Messages.Find(1, 10, byItemKeyCommentType(itemKey, commentType))
}

func GetAllMessages(page, size int) (records.Page, error) {
	return ctx.Messages.Find(page, size, op.Everything())
}

func SearchMessages(page, size int, param Message) (records.Page, error) {
	return ctx.Messages.Find(page, size, byExpression(param))
}

func UpdateNodeMessage(itemKey hsk.Key, commentType commenttype.Enum, data Message) error {
	rec, err := ctx.Messages.FindFirst(byItemKeyCommentType(itemKey, commentType))

	if err != nil {
		return err
	}

	return ctx.Messages.Update(rec.GetKey(), data)
}

func UpdateMessage(key hsk.Key, data Message) error {
	return ctx.Messages.Update(key, data)
}
