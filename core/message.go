package core

import (
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/husk"
)

type Message struct {
	UserKey     husk.Key
	ItemKey     husk.Key
	UpVotes     int64  `hsk:"null"`
	DownVotes   int64  `hsk:"null"`
	Text        string `hsk:"size(512)"`
	CommentType string
	Voters      []husk.Key
	Children    []Message
	UserImage   string //gravatar id
}

func (msg Message) Valid() error {
	return husk.ValidateStruct(&msg)
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

func GetMessageByKey(key husk.Key) (husk.Recorder, error) {
	return ctx.Messages.FindByKey(key)
}

func GetNodeMessage(itemKey husk.Key, commentType commenttype.Enum) (husk.Recorder, error) {
	return ctx.Messages.FindFirst(byItemKeyCommentType(itemKey, commentType))
}

func GetAllMessages(page, size int) (husk.Collection, error) {
	return ctx.Messages.Find(page, size, husk.Everything())
}

func SearchMessages(page, size int, param Message) (husk.Collection, error) {
	return ctx.Messages.Find(page, size, byExpression(param))
}

func UpdateNodeMessage(itemKey husk.Key, commentType commenttype.Enum, data Message) error {
	rec, err := ctx.Messages.FindFirst(byItemKeyCommentType(itemKey, commentType))

	if err != nil {
		return err
	}

	err = rec.Set(data)

	if err != nil {
		return err
	}

	err = ctx.Messages.Update(rec)

	if err != nil {
		return err
	}

	return ctx.Messages.Save()
}

func UpdateMessage(key husk.Key, data Message) error {
	rec, err := ctx.Messages.FindByKey(key)

	if err != nil {
		return err
	}

	err = rec.Set(data)

	if err != nil {
		return err
	}

	err = ctx.Messages.Update(rec)

	if err != nil {
		return err
	}

	return ctx.Messages.Save()
}
