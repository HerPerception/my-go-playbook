package main

import (
	"fmt"
	"os"
	"strings"
)

func ValidateText(text string) (rune, error) {
	for _, ch := range text {
		if ch < 32 || ch > 126 {
			return ch, fmt.Errorf("Input text contains unsupported character: %c", ch)
		}
	}
	return 0, nil
}

func GenerateArt(text, filename string) (string, error) {
	if _, err := ValidateText(text); err != nil {
		return "", err
	}
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	str := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	lines := strings.Split(strings.ReplaceAll(string(fileContent), "\r\n", "\n"), "\n")
	lines = lines[1:]
	var result []string

	for _, each := range str {
		for row := 0; row < 8; row++ {
			var builder strings.Builder
			for _, ch := range each {
				position := int(ch-32)*9 + row
				builder.WriteString(lines[position])
			}
			result = append(result, builder.String())
		}
	}
	return strings.Join(result, "\n"), nil
}
