package validations

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var largeDef = ""

func init() {
	bts, err := ioutil.ReadFile("large_def.json")
	if err != nil {
		fmt.Print(err)
	}
	largeDef = string(bts)
}

func TestContainerDefinition(t *testing.T) {
	var tests = []struct {
		name string
		definition string
		want bool
	}{
		{"shall_fail_", "a([]{}())", false},
		{"shall_sucess_", "[]{}()", true},
		{"shall_skip_other_chars_", "[a]{vcv}(11)", true},
		{"shall_sucess_", "([]{}())", true},
		{"shall_fail_", "([]{}()){", false},
		{"shall_fail_", "({)", false},
		{"shall_fail_", ")", false},
		{"shall_fail_", "", false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.name, tt.definition)
		t.Run(testname, func(t *testing.T) {
			ans := ContainerDefinition(tt.definition)
			if ans != tt.want {
				t.Errorf("test: %v, got %v, want %v",testname, ans, tt.want)
			}
		})
	}
}

func BenchmarkDefinition(b *testing.B) {
	for i := 0 ; i < 100 ; i++ {
		ContainerDefinition(largeDef)
	}
}

