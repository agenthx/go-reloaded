package piscine

import "regexp"

func Apostrophe(sent string) string {
	reAp:=regexp.MustCompile(`'\s*(\w*)\s*'`)
	return reAp.ReplaceAllString(sent, "'$1'")
}
