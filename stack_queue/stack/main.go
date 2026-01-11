package main

import (
	"fmt"
)

func main() {
	s := Stack{}
	fmt.Println(s.stack)
	s.push(1)
	fmt.Println(s.stack)
	s.push(2)
	fmt.Println(s.stack)
	fmt.Println(s.pop())
	fmt.Println(s.stack)
	fmt.Println(s.pop())
	fmt.Println(s.stack)
}

type Stack struct {
	stack []any
}

func (s *Stack) push(value any) {
	s.stack = append(s.stack, value)
}

func (s *Stack) pop() any {
	if len(s.stack) != 0 {
		value := s.stack[len(s.stack) - 1]
		s.stack = s.stack[:len(s.stack) - 1]
		return value
	}

	return nil
}