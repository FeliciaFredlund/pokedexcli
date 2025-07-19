package main

import "strings"

func cleanInput(text string) []string {
	words := []string{}

	text = strings.TrimSpace(strings.ToLower(text))
	textRemains := true

	for textRemains {
		var word string
		word, text, textRemains = strings.Cut(text, " ")
		words = append(words, word)
		text = strings.TrimSpace(text)
	}

	return words
}
