package main

import (
	"strings"
)

func cleanInput(text string) []string {
	textLower := strings.ToLower(text)
	textSplit := strings.Fields(textLower)
	return textSplit
}
