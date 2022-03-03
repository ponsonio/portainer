package util

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewRuneStack()
	//shall error
	_, errPop := s.Pop()

	if errPop == nil && errPop.Error() != "no data in the stack"{
		t.Errorf("shall give correct error, when empty")
	}
	s.Push('a')
	s.Push('b')
	s.Push('c')

	if s.Size() != 3 {
		t.Errorf("shall give correct size")
	}

	 c, errPopNil := s.Pop()

	if errPopNil != nil {
		t.Errorf("shall pop with no error")
	}

	if c != 'c' {
		t.Errorf("shall pop last rune")
	}

}
