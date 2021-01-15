package main

import (
	"bytes"
	"ekz/tools"
	"ekz/tree"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type App struct {
	client    http.Client
	baseURL   string
	parser    tree.Parser
	reader    tools.MyReader
	processor tree.Processor
}

func (a *App) process(data io.Reader) (io.Reader, error) {
	return a.processor.Process(data)
}

func (a *App) sendData(reader io.Reader) (*http.Response, error) {
	request, err := http.NewRequest("POST", a.baseURL, reader)
	if err != nil {
		return nil, err
	}
	res, err := a.client.Do(request)
	return res, nil
}

func (a *App) readData() (io.Reader, error) {
	buf := make([]byte, 0)
	buf, err := a.reader.Read()
	return bytes.NewReader(buf), err
}

func main() {
	client := &App{http.Client{},
		"http://localhost:8080/",
		tree.Parser{},
		&tools.FileReader{FileName: "cmd\\client\\file.json"},
		&tree.DoNothing{}}
	jsonReader, err := client.readData()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonReader, err = client.process(jsonReader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res, err := client.sendData(jsonReader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(bytes))
}
