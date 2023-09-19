package helpers

import (
	"io"
	"os"
)

func CopyFile(src, dst string) (err error) {
	fin, err := os.Open(src)
	if err != nil {
		return
	}
	defer fin.Close()

	fout, err := os.Create(dst)
	if err != nil {
		return
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)

	return
}

func MoveFile(src, dst string) (err error) {
	err = CopyFile(src, dst)
	if err != nil {
		return
	}

	err = os.Remove(src)

	return
}
