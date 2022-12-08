package ChangeData

import (
	"strings"
)

func Replacer(s string) string {
	replacer := strings.NewReplacer(" ", "%20")
	return replacer.Replace(s)
}
