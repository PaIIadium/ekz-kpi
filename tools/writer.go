package tools

import (
	"os"
)

type FileWriter struct {
	FileName string
}

func (fw *FileWriter) Write(data []byte) (int, error) {
	file, err := os.OpenFile(fw.FileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return 0, err
	}

	if err != nil {
		return 0, err
	}
	num, err := file.Write(data)
	return num, err
}
