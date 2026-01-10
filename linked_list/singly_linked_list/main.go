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

func (l *LinkedList) reverseIterative() {
	var previous *Node
	current := l.head
	for current != nil {
		next := current.next
		current.next = previous

		previous = current
		current = next
	}
	l.head = previous
}

func _reverseRecursive(current *Node, previous *Node) *Node {
	if current == nil {
		return previous
	}
	
	next := current.next
	current.next = previous
	return _reverseRecursive(next, current)
}

func (l *LinkedList) reverseRecursive() {
	l.head = _reverseRecursive(l.head, nil)
}

func _reverseEvenRecursive(head *Node, previous *Node) *Node {
	if head == nil {
		return nil
	}

	current := head
	for current != nil && int(current.data.(int)) % 2 == 0 {
		next := current.next
		current.next = previous

		previous = current
		current = next
	}

	if current != head {
		head.next = current
		// 再度別リストとして再帰関数を呼び出す
		_reverseEvenRecursive(current, nil)
		return previous
	} else {
		head.next = _reverseEvenRecursive(head.next, head)
		return head
	}
}

func (l *LinkedList) reverseEven() {
	l.head = _reverseEvenRecursive(l.head, nil)
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
	l.append(4)
	l.append(5)
	l.print()
	l.reverseIterative()
	l.print()
	l.reverseRecursive()
	l.print()
	l.append(2)
	l.append(4)
	l.append(6)
	l.append(8)
	l.append(9)
	l.append(10)
	l.append(12)
	l.reverseEven()
	l.print()
	// fmt.Println("LinkedList Head: ", l.head.data)
	// fmt.Println("LinkedList Next: ", l.head.next.data)
	// fmt.Println("LinkedList Next Next: ", l.head.next.next.data)
	// fmt.Println("LinkedList Next Next Next: ", l.head.next.next.next.data)
}
