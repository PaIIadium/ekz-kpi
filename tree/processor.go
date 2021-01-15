package tree

import (
	"io"
)

type Processor interface {
	Process(io.Reader) (io.Reader, error)
}

type DoNothing struct {
}

func (d *DoNothing) Process(jsonReader io.Reader) (io.Reader, error) {
	return jsonReader, nil
}

type MultiplyNodes struct {
	parser Parser
}

func (s *MultiplyNodes) RecursiveMultiply(root *Node, product int) int {
	for _, val := range root.Children {
		product *= s.RecursiveMultiply(val, 1)
	}
	return root.Value * product
}

func (s *MultiplyNodes) Process(jsonReader io.Reader) (io.Reader, error) {
	tree, err := s.parser.decode(jsonReader)
	if err != nil {
		return nil, err
	}
	queue := []*Node{tree.Root}
	for len(queue) > 0 {
		curNode := queue[0]
		product := s.RecursiveMultiply(curNode, 1)
		if len(curNode.Children) > 0 {
			for _, child := range curNode.Children {
				queue = append(queue, child)
			}
		}
		curNode.Value = product
		queue = queue[1:]
	}

	jsonReader, err = s.parser.encode(*tree)
	return jsonReader, err
}
