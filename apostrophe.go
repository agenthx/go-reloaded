package piscine

import "regexp"

func Apostrophe(sent string) string {
	reAp:=regexp.MustCompile(`\'\s*(.*?)\s*\'`)
	return reAp.ReplaceAllString(sent, "'$1'")
}