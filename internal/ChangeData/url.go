package ChangeData

import (
	"strings"
)

func Replacer(s string) string {
	replacer1 := strings.NewReplacer(" ", "%20")
	return replacer1.Replace(s)
}
