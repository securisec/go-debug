package debug

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
	"github.com/mitchellh/mapstructure"
)

const (
	FgRed      = color.FgRed
	FgBlue     = color.FgBlue
	FgYellow   = color.FgYellow
	FgGreen    = color.FgGreen
	FgCyan     = color.FgCyan
	FgMagenta  = color.FgMagenta
	FgWhite    = color.FgWhite
	FgBlack    = color.FgBlack
	Bold       = color.Bold
	Italic     = color.Italic
	Underline  = color.Underline
	BlinkRapid = color.BlinkRapid
	BlinkSlow  = color.BlinkSlow
	BgRed      = color.BgRed
	BgBlue     = color.BgBlue
	BgYellow   = color.BgYellow
	BgGreen    = color.BgGreen
	BgCyan     = color.BgCyan
	BgMagenta  = color.BgMagenta
	BgWhite    = color.BgWhite
	BgBlack    = color.BgBlack
)

type Config struct {
	Namespace string
	Style     []color.Attribute
	ShowInfo  bool
}

var ConfigDefault = Config{
	Namespace: "DEBUG",
	Style:     []color.Attribute{FgGreen},
	ShowInfo:  true,
}

func New(config ...Config) func(...interface{}) {
	cfg := ConfigDefault
	if len(config) > 0 {
		cfg = config[0]
		if cfg.Namespace == "" {
			cfg.Namespace = ""
		}
	}
	formatter := color.New(cfg.Style...)
	return func(data ...interface{}) {
		var (
			l   interface{}
			out interface{}
		)

		s := make([]string, 0)
		for _, d := range data {
			s = append(s, fmt.Sprintf("%v", d))
		}

		err := mapstructure.Decode(out, l)
		if err != nil {
			l = s
		}

		if cfg.ShowInfo {
			callingFunction := ""
			pc, file, no, ok := runtime.Caller(1)
			details := runtime.FuncForPC(pc)
			if ok {
				callingFunction = fmt.Sprintf("%s#%s:%d\n", file, details.Name(), no)
			}
			formatter.Printf("%s %s%s\n", cfg.Namespace, callingFunction, l)
		} else {
			formatter.Printf("%s %s\n", cfg.Namespace, l)
		}
	}
}
