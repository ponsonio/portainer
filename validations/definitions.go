package validations

import (
	"github.com/jcabrera/portainer/util"
)

var (
	brackets = map[rune]bool{
		'}':true,
		')':true,
		']':true,
		'{':true,
		'(':true,
		'[':true,
	}
	openings = map[rune]bool{
		'{': true,
		'(': true,
		'[': true,
	}

	closures = map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
)

func ContainerDefinition(definition string) bool {
	stack := util.NewRuneStack()
	if len(definition) == 0 {
		return false
	}

	//shall start with opening
	_ , st := openings[rune(definition[0])]
	if !st {
		return false
	}

	for _, c := range definition {
		_ , t := brackets[c]
		//char I don't care
		if !t {
			continue
		}
		//if it's opening
		_ , e := openings[c]
		if e {
			stack.Push(c)
			continue
		}
		//needs to be closure
		expect , _ := closures[c]
		match, sErr := stack.Pop()
		if sErr != nil {
			if sErr.Error() != "no data in the stack" {
				panic("no expected stack error")
			}
			return false
		}
		if match != expect {
			return false
		}
	}
	if stack.Size() > 0 {
		return false
	}
	return true
}