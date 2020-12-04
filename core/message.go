package core

import (
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/husk/validation"
)

type Message struct {
	SubjectID   string //oidc user
	ItemKey     hsk.Key
	UpVotes     int64  `hsk:"null"`
	DownVotes   int64  `hsk:"null"`
	Text        string `hsk:"size(512)"`
	CommentType commenttype.Enum
	Voters      []string
	Children    []Message
	UserImage   string //gravatar id
}

func (msg Message) Valid() error {
	return validation.Struct(msg)
}

func (msg Message) SubmitMessage() error {
	msg.UpVotes = 0
	msg.DownVotes = 0

	_, err := ctx.Messages.Create(msg)

	if err != nil {
		return err
	}

	return ctx.Messages.Save()
}

func GetMessageByKey(key hsk.Key) (hsk.Record, error) {
	return ctx.Messages.FindByKey(key)
}

func GetNodeMessage(itemKey hsk.Key, commentType commenttype.Enum) (hsk.Record, error) {
	return ctx.Messages.FindFirst(byItemKeyCommentType(itemKey, commentType))
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
