package util

import "fmt"

/**
type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s *Stack[T]) Size() int {
	return len(s.vals)
}
*/

type RuneStack interface {
	Push(r rune)
	Pop() (rune, error)
	Size() int
}

type simpleRuneStack struct {
	data  []rune
}

func (s *simpleRuneStack) Push(r rune) {
	s.data = append(s.data, r)
}

func (s *simpleRuneStack) Pop() (rune, error) {
	if len(s.data) > 0 {
		r := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return r, nil
	}
	return 0, fmt.Errorf("no data in the stack")
}

func (s *simpleRuneStack)  Size() int {
	return len(s.data)
}

func NewRuneStack() RuneStack {
	return &simpleRuneStack{}
}
