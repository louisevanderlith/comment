package core_test

import (
	"testing"

	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/comment/core/commenttype"
)

func init() {
	core.CreateContext()
}

func TestSubmitMessage_AllEmpty_Invalid(t *testing.T) {
	msg := core.Message{}
	err := msg.SubmitMessage()

	if err == nil {
		t.Error("Expecting validation errors.")
	}
}

func TestSubmitMessage_TextEmpty_Invalid(t *testing.T) {
	msg := core.Message{}
	msg.CommentType = commenttype.Stock

	err := msg.SubmitMessage()

	if err == nil {
		t.Error("Expecting 'Text cant be empty'")
	}
}

func TestSubmitMessage_CommentTypeEmpty_Invalid(t *testing.T) {
	msg := core.Message{}
	msg.Text = "Testing some message"

	err := msg.SubmitMessage()

	if err == nil {
		t.Error("Expecting 'CommentType cant be empty'")
	}
}

func TestSubmitMessage_RequiredOnly_Valid(t *testing.T) {
	msg := core.Message{}
	msg.CommentType = commenttype.Stock
	msg.Text = "Testing some message"
	msg.UserImage = "jshdkfjha23,mnsdflkjx!"

	err := msg.SubmitMessage()

	if err != nil {
		t.Error(err)
	}
}
