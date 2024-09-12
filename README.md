# gohelper

[![Go](https://github.com/imajinyun/gohelper/actions/workflows/go.yml/badge.svg)](https://github.com/imajinyun/gohelper/actions/workflows/go.yml)

A package of Go helper functions for business development.

## Installation

Make sure that Go is installed on your computer. Type the following command in your terminal:

```bash
go get github.com/imajinyun/gohelper
```

Add following line in your `*.go` file:

```go
import "github.com/imajinyun/gohelper"
```

## Examples

### Cond

1. Simplify the judgment of if else to a single line of code:

```go
flag, expr := 0, "some version"
if expr == "v1" {
  flag = 1
} else if expr == "v2" {
  flag = 2
}

flag = gohelper.If(expr == "v1", 1, gohelper.If(expr == "v2", 2, 0))
```

### Conv

1. Convert data to string:

```go
// Output: [hello world]
gohelper.ToString([]string{"hello", "world"})
```

2. Convert data to JSON string:

```go
// Output: {"id":1,"name":"jack"}
gohelper.ToJson(map[string]any{"id": 1, "name": "jack"})
```

### Date

1. Get the current date and time object:

```go
// Output: 2024-09-12 12:02:15
tim, err := gohelper.NowDateTime("Asia/Shanghai", "2006-01-02 15:04:05")
if err != nil {
  panic(err)
}
tim.ToString()
```

2. Get the year, month, day, hour, minute, and second of the date and time object:

```go
// Output: 2024
tim.Year()

// Output: September
tim.Month()

// Output: 12
tim.Day()

// Output: 12
tim.Hour()

// Output: 2
tim.Minute()

// Output: 15
tim.Second()
```

### Map

1. If the specified key cannot obtain a value, return the given default value:

```go
// Output: default value
mps := make(map[string]any)
mps.GetOrDefault("mykey", "default value")
```

### Str

####

1. Generate string with options (include uppercase, numbers, and symbols):

```go
// Output: knvmfcmpfqiqcbrh
gohelper.RandStr(16)

// Output: nD>fKDvaF\R+1h.G
gohelper.RandStrWithOption(16, Option{
  IncludeNumber:    true,
  IncludeUppercase: true,
  IncludeSymbol:    true,
})
```
