# go-debug

<p align="center">
    <img src="https://i.imgur.com/L1EO89a.png" width="80%">
</p>

`go-debug` is a similar library to [npm debug package](https://github.com/visionmedia/debug). This module was written mostly for personal consumption and does not offer anywhere near the amount of customization offered by the npm package.

`go-debug` uses the `DEBUG` envrionment variable to log data. This envar can contain multiple comma seperated values indicating different namespaces. If a `*` is used as the value of this envar, it will log all namespaces. If the envar does not have any value set, all logs will be ignored.

General rules to remember when setting the `DEBUG` environment:
- Multiple namespaces can be specified with a `,`
- `*` Will log all
- `TAG` Will log tag only
- `!TAG` Will not be logged

Example:
`export DEBUG=APP`
```go
debugMe := gdebug.New()
debugMe("some log")
// This will not log

newDebug := debug.New(debug.Config{
    Namespace: "APP",
    Style: []color.Attribute{
        Yellow,
        Bold,
    },
})
newDebug("something", 1, true)
// This will log
```

## Install
```bash
go get -u github.com/securisec/go-debug
```

## Example

```go
package main

import (
	"github.com/fatih/color"
	gdebug "github.com/securisec/go-debug"
)

func main() {
	debugApp := gdebug.New()
	debugApp("Hello")

	debugPretty := gdebug.New(gdebug.Config{
		Namespace: "PRETTY",
		Style:     []color.Attribute{gdebug.Blue},
		Pretty:    true,
	})
	t := &Test{
		Some:    "some",
		Another: true,
		Data: Data{
			Something: "something",
		},
	}
	debugPretty(t)

	debugError := gdebug.New(gdebug.Config{Namespace: "ERROR", Style: []color.Attribute{gdebug.Red}})
	debugError(t)
	debugWarning := gdebug.New(gdebug.Config{Namespace: "WARNING", Style: []color.Attribute{gdebug.Yellow}})
	debugWarning("Warning", "message")
}

type Test struct {
	Some    string
	Another bool
	Data    Data
}

type Data struct {
	Something string
}

```