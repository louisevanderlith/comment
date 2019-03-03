package core_test

import (
	"testing"

	comment "github.com/louisevanderlith/comment/core"
)

func TestSubmitMessage_AllEmpty_Invalid(t *testing.T) {
	msg := comment.Message{}
	set := comment.SubmitMessage(msg)

	if set.Error == nil {
		t.Error("Expecting validation errors.")
	}
}

func TestSubmitMessage_TextEmpty_Invalid(t *testing.T) {
	msg := comment.Message{}
	msg.CommentType = comment.Advert

	set := comment.SubmitMessage(msg)

	if set.Error == nil {
		t.Error("Expecting 'Text cant be empty'")
	}
}

func TestSubmitMessage_CommentTypeEmpty_Invalid(t *testing.T) {
	msg := comment.Message{}
	msg.Text = "Testing some message"

	set := comment.SubmitMessage(msg)

	if set.Error == nil {
		t.Error("Expecting 'CommentType cant be empty'")
	}
}

func TestSubmitMessage_RequiredOnly_Valid(t *testing.T) {
	msg := comment.Message{}
	msg.CommentType = comment.Advert
	msg.Text = "Testing some message"

	set := comment.SubmitMessage(msg)

	if set.Error != nil {
		t.Error(set)
	}
}
