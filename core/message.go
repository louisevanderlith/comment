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
	Voters      map[husk.Key]struct{}
	Children    []Message
	UserImage   string //gravatar id
}

func (o Message) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func SubmitMessage(msg Message) husk.CreateSet {
	msg.UpVotes = 0
	msg.DownVotes = 0

	defer ctx.Messages.Save()

	return ctx.Messages.Create(msg)
}

func GetMessageByKey(key husk.Key) (husk.Recorder, error) {
	return ctx.Messages.FindByKey(key)
}

func GetMessage(itemKey husk.Key, commentType commenttype.Enum) (husk.Recorder, error) {
	return ctx.Messages.FindFirst(byItemKeyCommentType(itemKey, commentType))
}

func GetAllMessages(page, size int) husk.Collection {
	return ctx.Messages.Find(page, size, husk.Everything())
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
