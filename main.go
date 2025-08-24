package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	cleaned := []string{}

	lowerText := strings.ToLower(text)
	trimmedText := strings.TrimSpace(lowerText)
	for word := range strings.SplitSeq(trimmedText, " ") {
		if word != "" {
			cleaned = append(cleaned, word)
		}
	}

	return cleaned
}

func main() {
	fmt.Println("Hello, World!")
}
