package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordFrequency(input string) map[string]int {
	input = strings.ToLower(input)

	var cleanedInput strings.Builder
	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsSpace(char) {
			cleanedInput.WriteRune(char)
		}
	}

	words := strings.Fields(cleanedInput.String())

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	input := "Hello, hello! How are you? Are you doing well, well?"
	
	frequency := wordFrequency(input)
	
	fmt.Println("Word Frequency Count:")
	for word, count := range frequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}
