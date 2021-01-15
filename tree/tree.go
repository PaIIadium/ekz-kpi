package tree

type Node struct {
	Value    int
	Children []*Node
}

type Tree struct {
	Root *Node
}
