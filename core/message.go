package core

import (
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/husk/hsk"
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
