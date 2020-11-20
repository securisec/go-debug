package gdebug

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"github.com/fatih/color"
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
	// Out controls where something is being logged.
	// Optional. Default: os.Stderr
	Out io.Writer
	// Pretty print objects
	// Optional. Default: false
	Pretty bool
}

// ConfigDefault default config. Sets foreground color to green and shows additional info. DEBUG namespace.
var ConfigDefault = Config{
	Namespace: "DEBUG",
	Style:     []color.Attribute{color.FgCyan},
	ShowInfo:  false,
	Out:       os.Stderr,
	Pretty:    false,
}

// New create a new debug function
// Example:
// debugApp := gdebug.New()
// debugApp("Hello world")
func New(config ...Config) func(...interface{}) {
	cfg := ConfigDefault
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.Out == nil {
		cfg.Out = os.Stderr
	}
	formatter := color.New(cfg.Style...)
	nameSpace := checkDebugEnv(cfg.Namespace)
	return func(data ...interface{}) {
		var (
			hold interface{}
			err  error
		)

		check := make([]interface{}, 0)
		for _, d := range data {
			if ok := isErrorType(d); ok {
				check = append(check, d.(error).Error())
			}
			check = append(check, d)
		}

		if cfg.Pretty {
			hold, err = json.MarshalIndent(check, "", "  ")
		} else {
			hold, err = json.Marshal(check)
		}

		if err != nil {
			hold = data
		} else {
			hold = unsafeString(hold.([]byte))
		}

		if cfg.ShowInfo {
			callingFunction := ""
			pc, file, no, ok := runtime.Caller(1)
			details := runtime.FuncForPC(pc)
			if ok {
				callingFunction = fmt.Sprintf("%s#%s:%d\n", file, details.Name(), no)
			}
			if nameSpace {
				formatter.Printf("%s %s%s\n", cfg.Namespace, callingFunction, hold)
			}
		} else {
			if nameSpace {
				formatter.Printf("%s %s\n", cfg.Namespace, hold)
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
		if strings.HasPrefix(namespace, "!") && strings.Replace(namespace, "!", "", 1) == namespace {
			return false
		}
		if n == namespace {
			return true
		}
	}
	return false
}

func unsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func isErrorType(err interface{}) bool {
	t := reflect.TypeOf(err).String()
	if t == "*errors.errorString" {
		return true
	}
	return false
}
