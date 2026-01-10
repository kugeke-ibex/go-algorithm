package main

import (
	"fmt"
)

func main() {
	d := DoublyLinkedList{}
	d.append(1)
	d.append(2)
	d.append(3)
	d.insert(0)
	d.append(4)
	d.print()
	d.remove(2)
	d.print()
	d.remove(0)
	d.print()
	d.remove(4)
	d.print()
	// fmt.Println(d.head.data)
	// fmt.Println(d.head.next.data)
	// fmt.Println(d.head.next.next.data)
	// fmt.Println(d.head.next.next.next.data)
	d.append(5)
	d.append(7)
	d.print()
	d.reverseIterative()
	d.print()
	d.reverseRecursive()
	d.print()
	d.reverseRecursive()
	d.print()
	d.sort()
	d.print()
}

type Node struct {
	data any
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
}

func (d *DoublyLinkedList) append(data any) {
	newNode := &Node{data: data, next:nil, prev:nil}
	if d.head == nil {
		d.head = newNode
		return
	}

	current := d.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	newNode.prev = current
}

func (d *DoublyLinkedList) insert(data any) {
	newNode := &Node{data: data, next:nil, prev:nil}

	if d.head == nil {
		d.head = newNode
		return
	}

	current := d.head
	current.prev = newNode
	newNode.next = current
	d.head = newNode
}

func (d DoublyLinkedList) print() {

	current := d.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println("--------------------------------")
}

func (d *DoublyLinkedList) remove(target any) {
	
	var previous *Node
	current := d.head
	for current != nil {
		if current.data == target {
			if previous == nil {
				d.head = current.next
			} else {
				previous.next = current.next
				if current.next != nil {
					current.next.prev = previous
				}
			}
			return 
		}

		previous = current
		current = current.next
	}
}

func (d *DoublyLinkedList) reverseIterative() {
	var previous *Node
	current := d.head
	for current != nil {
		next := current.next
		current.next = previous
		current.prev = next

		previous = current
		current = next
	}
	d.head = previous
}

func _reverseRecursive(current *Node, previous *Node) *Node {
	if current == nil {
		return previous
	}

	next := current.next
	current.prev = next
	current.next = previous

	return _reverseRecursive(next, current)
}

func (d *DoublyLinkedList) reverseRecursive() {
	d.head = _reverseRecursive(d.head, nil)
}

func (d *DoublyLinkedList) sort() {
	if d.head == nil {
		return
	}

	current := d.head
	for current.next != nil {
		next := current.next
		for next != nil {
			if current.data.(int) > next.data.(int) {
				current.data, next.data = next.data, current.data
			}
			next = next.next
		}
		current = current.next
	}
}
