package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"strings"

	"github.com/louisevanderlith/comment/core/commenttype"
)

type messageFilter func(obj Message) bool

func (f messageFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Message))
}

func byItemKeyCommentType(itemKey hsk.Key, ct commenttype.Enum) messageFilter {
	return func(obj Message) bool {
		return obj.ItemKey.Compare(itemKey) == 0 && obj.CommentType == ct
	}
}

//will look for 'containing' text, userKeys, & commentTypes
func byExpression(param Message) messageFilter {
	return func(obj Message) bool {
		hasText := false
		hasUser := false
		hasType := false

		if len(param.Text) > 0 {
			hasText = strings.Contains(obj.Text, param.Text)
		}

		if param.SubjectID != "" {
			hasUser = obj.SubjectID == param.SubjectID
		}

		if param.CommentType > 0 {
			hasType = obj.CommentType == param.CommentType
		}

		return hasText || hasUser || hasType
	}
}
