package main

import (
	"fmt"
	"io"
	"net/http"
)

func FetchURL(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	return string(body), err
}
func main() {
	fmt.Println(FetchURL("https://httpbin.org/get"))
}
