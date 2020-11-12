# go-debug

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
debugMe := debug.New()
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