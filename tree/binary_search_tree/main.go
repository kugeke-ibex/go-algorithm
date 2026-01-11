package main

import (
	"fmt"
)

func main() {
	var root *Node
	root = root.Insert(root, 3)
	root = root.Insert(root, 6)
	root = root.Insert(root, 5)
	root = root.Insert(root, 7)
	root = root.Insert(root, 1)
	root = root.Insert(root, 10)
	root = root.Insert(root, 2)
	fmt.Println(root.left.value)
	fmt.Println(root.left.right.value)
}

type Node struct {
	value int
	left *Node
	right *Node
}

func (n *Node) Insert(node *Node, value int) *Node{
	if node == nil {
		return &Node{value: value}
	}

	if value < node.value {
		node.left = n.Insert(node.left, value)
	} else {
		node.right = n.Insert(node.right, value)
	}

	return node
}

