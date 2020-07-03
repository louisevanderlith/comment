package core

import (
	"strings"

	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/husk"
)

type messageFilter func(obj Message) bool

func (f messageFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(Message))
}

func byItemKeyCommentType(itemKey husk.Key, commentType commenttype.Enum) messageFilter {
	typeStr := commentType.String()
	return func(obj Message) bool {
		return obj.ItemKey == itemKey && obj.CommentType == typeStr
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

		if param.UserKey != husk.CrazyKey() {
			hasUser = obj.UserKey == param.UserKey
		}

		if len(param.CommentType) > 0 {
			hasType = obj.CommentType == param.CommentType
		}

		return hasText || hasUser || hasType
	}
}
