package main

import (
	"fmt"
	"os"
)

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("ReadFile %s: %w", path, err)
	}
	return string(data), nil
}

func TaskFive() {
	content, err := ReadFile("files.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(content)
}
