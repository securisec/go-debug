package gdebug

import (
	"errors"
	"os"
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
			31,
			Bold,
		},
		ShowInfo: true,
	})
	s := &SomeStruct{
		A: "a",
		B: 1,
		C: true,
		D: map[string]interface{}{"some": "data", "another": 1},
	}
	debugMe(s)
}

func TestJSON(t *testing.T) {
	debugMe := New(Config{Pretty: true})
	s := &SomeStruct{A: "lalla"}
	debugMe(s)
}

func TestIO(t *testing.T) {
	debugMe := New(Config{
		Out: os.Stdout,
	})
	debugMe("stdout")
}

func TestCoverage(t *testing.T) {
	debugMe := New(Config{})
	debugMe("test")
}

func TestError(t *testing.T) {
	debugMe := New()
	err := errors.New("some error")
	debugMe(err)
}

type SomeStruct struct {
	A string
	B int `json:"b"`
	C bool
	D interface{}
}
