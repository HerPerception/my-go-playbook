package main

import (
	"io"
	"os"
)

func CopyFile(src, dst string) error {
	file, err := os.OpenFile(src, os.O_RDWR, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	dstfile, err := os.Create(dst)
	defer dstfile.Close()

	if err != nil {
		return err
	}
	_, err = io.Copy(dstfile, file)
	if err != nil {
		return err
	}
	return nil
}
func main() {

}
