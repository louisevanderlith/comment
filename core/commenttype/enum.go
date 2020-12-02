package commenttype

import (
	"strings"
)

type Enum = int

const (
	Profile Enum = iota
	Stock
	Article
	Child
)

var vals = [...]string{
	"Profile",
	"Stock",
	"Article",
	"Child",
}

func StringEnum(r Enum) string {
	return vals[r]
}

func GetEnum(name string) Enum {
	var result Enum

	for k, v := range vals {
		if strings.ToUpper(name) == strings.ToUpper(v) {
			return k
		}
	}

	return result
}
