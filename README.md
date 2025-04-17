# gokit

[![Go](https://github.com/imajinyun/gokit/actions/workflows/go.yml/badge.svg)](https://github.com/imajinyun/gokit/actions/workflows/go.yml)

A package of Go kit functions for business development.

## Installation

Make sure that Go is installed on your computer. Type the following command in your terminal:

```bash
go get github.com/imajinyun/gokit
```

Add following line in your `*.go` file:

```go
import "github.com/imajinyun/gokit"
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

flag = gokit.If(expr == "v1", 1, gokit.If(expr == "v2", 2, 0))
```

### Conv

1. Convert data to string:

```go
// Output: [hello world]
gokit.ToString([]string{"hello", "world"})
```

2. Convert data to JSON string:

```go
// Output: {"id":1,"name":"jack"}
gokit.ToJson(map[string]any{"id": 1, "name": "jack"})
```

### Date

1. Get the current date and time object:

```go
// Output: 2024-09-12 12:02:15
tim, err := gokit.NowDateTime("Asia/Shanghai", "2006-01-02 15:04:05")
if err != nil {
  panic(err)
}
tim.ToString()
```

2. Get the year, month, day, hour, minute, and second of the date and time object:

```go
// Output: 2024
tim.Year()

// Output: 256
tim.YearDay()

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

3. Get the begin and end of the day of the date and time object:

```go
// Output: 2024-09-12 00:00:00
tim.BeginOfDay().Format(time.DateTime)

// Output: 2024-09-12 23:59:59
tim.BeginOfDay().Format(time.DateTime)

// Output: 2024-09-01 00:00:00
tim.BeginOfMonth().Format(time.DateTime)

// Output: 2024-09-01 23:59:59
tim.EndOfMonth().Format(time.DateTime)

// Output: 2024-01-01 00:00:00
tim.BeginOfYear().Format(time.DateTime)

// Output: 2024-12-31 23:59:59
tim.EndOfYear().Format(time.DateTime)
```

### Map

1. If the specified key cannot obtain a value, return the given default value:

```go
// Output: default value
mps := make(map[string]any)
GetOrDefault("mykey", "default value")
```

### Str

####

1. Generate string with options (include uppercase, numbers, and symbols):

```go
// Output: knvmfcmpfqiqcbrh
gokit.RandStr(16)

// Output: nD>fKDvaF\R+1h.G
gokit.RandStrWithOption(16, Option{
  IncludeNumber:    true,
  IncludeUppercase: true,
  IncludeSymbol:    true,
})
```
