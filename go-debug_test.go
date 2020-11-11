package debug

import (
	"testing"

	"github.com/fatih/color"
)

func TestDefault(t *testing.T) {
	debugMe := New()
	debugMe("test string")
}

func TestCustom(t *testing.T) {
	debugMe := New(Config{
		Namespace: "TEST",
		Style: []color.Attribute{
			Green,
			Bold,
		},
		ShowInfo: false,
	})
	s := &SomeStruct{
		A: "a",
		B: 1,
		C: true,
		D: map[string]interface{}{"some": "data", "another": 1},
	}
	debugMe(s)
}

func TestCoverage(t *testing.T) {
	debugMe := New(Config{})
	debugMe("test")
}

type SomeStruct struct {
	A string
	B int `json:"b"`
	C bool
	D interface{}
}
