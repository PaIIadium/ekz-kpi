package main

import (
	"ekz/tools"
	"ekz/tree"
	"io"
	"io/ioutil"
	"net/http"
)

// HTTP handler.
type HTTPHandlerFunc http.HandlerFunc

// httpHandler creates a new instance of channels HTTP handler.
func httpHandler(processor tree.Processor, writer io.Writer) HTTPHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handleReceiveTree(r, rw, processor, writer)
		}
	}
}

func handleReceiveTree(r *http.Request, rw http.ResponseWriter, processor tree.Processor, writer io.Writer) {
	data, err := processor.Process(r.Body)
	if err != nil {
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	bytes, err := ioutil.ReadAll(data)
	if err != nil {
		tools.WriteJsonBadRequest(rw, "File did not save")
		return
	}
	_, err = writer.Write(bytes)
	if err == nil {
		tools.WriteJsonOk(rw, "Successfully received and saved")
		return
	}
	tools.WriteJsonBadRequest(rw, "File did not save")
}
