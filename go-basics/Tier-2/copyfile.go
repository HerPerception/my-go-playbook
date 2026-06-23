package main

import (
	"io"
	"os"
)

func CopyFile(src, dst string) error {
	file, err := os.OpenFile(src, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	dstfile, err := os.Create(dst)

	if err != nil {
		return err
	}
	defer dstfile.Close()

	_, err = io.Copy(dstfile, file)
	if err != nil {
		return err
	}
	return nil
}

// func main() {

// }
