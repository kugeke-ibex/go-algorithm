package main

import (
	"fmt"
)

type Node struct {
	data any
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) append(data any) {
	newNode := Node{data: data, next: nil}
	if l.head == nil {
		l.head = &newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &newNode
}

func (l *LinkedList) insert(data any) {
	newNode := Node{data: data, next: nil}
	newNode.next = l.head
	l.head = &newNode
}

func (l LinkedList) print() {
	current :=  l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println("--------------------------------")
}

func (l *LinkedList) remove(target any) {
	current := l.head
	var previous *Node
	for current != nil {
		if current.data == target {
			if previous == nil {
				l.head = current.next
			} else {
				previous.next = current.next
			}
			return
		}
		previous = current
		current = current.next
	}
}

func main() {
	l := LinkedList{}
	l.append(1)
	l.append(2)
	l.append(3)
	l.insert(0)
	l.print()
	l.remove(2)
	l.print()
	l.remove(0)
	l.print()
	l.append(3)
	l.print()
	l.remove(3)
	l.print()
	// fmt.Println("LinkedList Head: ", l.head.data)
	// fmt.Println("LinkedList Next: ", l.head.next.data)
	// fmt.Println("LinkedList Next Next: ", l.head.next.next.data)
	// fmt.Println("LinkedList Next Next Next: ", l.head.next.next.next.data)
}
