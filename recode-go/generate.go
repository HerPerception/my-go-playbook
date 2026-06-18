package main

import (
	"fmt"
	"os"
	"strings"
)

func Generator(text, banner string) (string, error) {
	data, err := os.ReadFile(banner)
	if err != nil {
		return "", fmt.Errorf("error with file")
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	lines = lines[1:]
	str := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	var result []string
	for _, each := range str {
		for row := 0; row < 8; row++ {
			var builder strings.Builder
			for _, ch := range each {
				pos := int(ch-32)*9 + row
				builder.WriteString(lines[pos])
			}
			result = append(result, builder.String())
		}
	}
	return strings.Join(result, "\n"), nil
}
