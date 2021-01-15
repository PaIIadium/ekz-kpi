package tree

import (
	"bytes"
	"encoding/json"
	"io"
)

type Parser struct {
}

func (p *Parser) decode(jsonReader io.Reader) (*Tree, error) {
	var t Tree
	err := json.NewDecoder(jsonReader).Decode(&t)
	return &t, err
}

func (p *Parser) encode(tree Tree) (io.Reader, error) {
	js, err := json.Marshal(tree)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(js), nil
}
