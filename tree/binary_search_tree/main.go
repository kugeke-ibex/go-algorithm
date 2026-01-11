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

	fmt.Println("--------------------------------")
	root.InOrder()
	fmt.Println("--------------------------------")
	root.PreOrder()
	fmt.Println("--------------------------------")
	root.PostOrder()
	fmt.Println("--------------------------------")

	fmt.Printf("Search 10: %v\n", root.Search(10))
	fmt.Printf("Search 5: %v\n", root.Search(5))
	fmt.Printf("Search 9: %v\n", root.Search(9))
	fmt.Printf("Search 4: %v\n", root.Search(4))

	fmt.Println("--------------------------------")
	root = root.Remove(6)
	root.InOrder()
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Insert(node *Node, value int) *Node {
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

func (n *Node) InOrder() {
	if n != nil {
		n.left.InOrder()
		fmt.Println(n.value)
		n.right.InOrder()
	}
}

func (n *Node) PreOrder() {
	if n != nil {
		fmt.Println(n.value)
		n.left.PreOrder()
		n.right.PreOrder()
	}
}

func (n *Node) PostOrder() {
	if n != nil {
		n.left.PostOrder()
		n.right.PostOrder()
		fmt.Println(n.value)
	}
}

func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}
	if value == n.value {
		return true
	} else if value < n.value {
		return n.left.Search(value)
	} else {
		return n.right.Search(value)
	}
}

func (n *Node) minNode() *Node {
	current := n
	for current.left != nil {
		current = current.left
	}

	return current
}

func (n *Node) Remove(value int) *Node {
	if n == nil {
		return nil
	}

	if value < n.value {
		n.left = n.left.Remove(value)
	} else if value > n.value {
		n.right = n.right.Remove(value)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		temp := n.right.minNode()
		n.value = temp.value
		n.right = n.right.Remove(temp.value)
	}

	return n
}
