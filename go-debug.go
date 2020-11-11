package debug

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/mitchellh/mapstructure"
)

const (
	// Red foreground
	Red = color.FgRed
	// Blue Blue foreground
	Blue = color.FgBlue
	// Yellow Yellow foreground
	Yellow = color.FgYellow
	// Green Green foreground
	Green = color.FgGreen
	// Cyan foreground
	Cyan = color.FgCyan
	// Magenta foreground
	Magenta = color.FgMagenta
	// White foreground
	White = color.FgWhite
	// Black foreground
	Black = color.FgBlack
	// Bold text
	Bold = color.Bold
	// Italic text
	Italic = color.Italic
	// Underline text
	Underline = color.Underline
	// BlinkRapid text
	BlinkRapid = color.BlinkRapid
	// BlinkSlow text
	BlinkSlow = color.BlinkSlow
	// BgRed background
	BgRed = color.BgRed
	// BgBlue background
	BgBlue = color.BgBlue
	// BgYellow background
	BgYellow = color.BgYellow
	// BgGreen background
	BgGreen = color.BgGreen
	// BgCyan background
	BgCyan = color.BgCyan
	// BgMagenta background
	BgMagenta = color.BgMagenta
	// BgWhite background
	BgWhite = color.BgWhite
	// BgBlack background
	BgBlack = color.BgBlack
)

// Config debug options
type Config struct {
	// Namespace Namespace to identify different debug statements
	// Optional. Default: DEBUG
	Namespace string
	// Style Array of styles for debug logging
	// Optional. Default: []color.Attribute{Green}
	Style []color.Attribute
	// ShowInfo Show additional information like file name, line number and function name
	// Optional. Default: false
	ShowInfo bool
}

// ConfigDefault default config. Sets foreground color to green and shows additional info. DEBUG namespace.
var ConfigDefault = Config{
	Namespace: "DEBUG",
	Style:     []color.Attribute{Green},
	ShowInfo:  false,
}

// New create a new debug function
// Example:
// debugApp := debug.New()
// debugApp("Hello world")
func New(config ...Config) func(...interface{}) {
	cfg := ConfigDefault
	if len(config) > 0 {
		cfg = config[0]
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

		l, err := json.Marshal(out)
		if err != nil {
			err = mapstructure.Decode(out, &l)
			if err != nil {
				l = s
			}
		}

		if _, ok := reflect.ValueOf(l).Interface().([]byte); ok {
			l = string(l.([]byte))
		}

		if cfg.ShowInfo {
			callingFunction := ""
			pc, file, no, ok := runtime.Caller(1)
			details := runtime.FuncForPC(pc)
			if ok {
				callingFunction = fmt.Sprintf("%s#%s:%d\n", file, details.Name(), no)
			}
			if checkDebugEnv(cfg.Namespace) {
				formatter.Printf("%s %s%s\n", cfg.Namespace, callingFunction, l)
			}
		} else {
			if checkDebugEnv(cfg.Namespace) {
				formatter.Printf("%s %s\n", cfg.Namespace, l)
			}
		}
	}
}

func checkDebugEnv(n string) bool {
	nameSpaces, ok := os.LookupEnv("DEBUG")
	if !ok {
		return false
	}
	if nameSpaces == "" {
		return false
	}
	if nameSpaces == "*" {
		return true
	}
	for _, namespace := range strings.Split(nameSpaces, ",") {
		namespace = strings.Trim(namespace, " ")
		if n == namespace {
			return true
		}
	}
	return false
}
