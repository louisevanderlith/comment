package commenttype

import (
	"strings"
)

type Enum int

const (
	Profile Enum = iota
	Stock
	Article
	Child
)

var commentTypes = [...]string{
	"Profile",
	"Stock",
	"Article",
	"Child",
}

func (r Enum) String() string {
	return commentTypes[r]
}

func GetEnum(name string) Enum {
	var result Enum

	for k, v := range commentTypes {
		if strings.ToUpper(name) == strings.ToUpper(v) {
			result = Enum(k)
			break
		}
	}

	return result
}
