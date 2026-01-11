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

	j := "{'key1': 'value1', 'key2': [1, 2, 3], 'key3': {'key4': 'value4'}}"
	fmt.Printf("Is valid format: %t\n", validateFormat(j))
	j = "{'key1': 'value1', 'key2': [1, 2, 3], 'key3': {'key4': 'value4'}}]"
	fmt.Printf("Is valid format: %t\n", validateFormat(j))
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

func validateFormat(chars string) bool {
	lookup := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	lookupValues := map[string]bool{}
	for _, value := range lookup {
		lookupValues[value] = true
	}

	s := Stack{}
	for _, char := range chars {
		if _, ok := lookup[string(char)]; ok {
			s.push(lookup[string(char)])
			continue
		}
		
		if _, ok := lookupValues[string(char)]; ok {
			if len(s.stack) == 0 {
				return false
			}
			if s.pop() != string(char) {
				return false
			}
			continue
		}
	}

	return len(s.stack) == 0
}