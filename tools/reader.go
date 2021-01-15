package tools

import (
	"io/ioutil"
	"os"
)

type MyReader interface {
	Read() ([]byte, error)
}

type FileReader struct {
	FileName string
}

func (fr *FileReader) Read() ([]byte, error) {
	file, err := os.OpenFile(fr.FileName, os.O_RDWR, 0755)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	return data, err
}
