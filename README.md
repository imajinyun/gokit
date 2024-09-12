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

```go
// Simplify the judgment of if else to a single line of code:
flag, expr := 0, "some version"
if expr == "v1" {
  flag = 1
} else if expr == "v2" {
  flag = 2
}

flag = gohelper.If(expr == "v1", 1, gohelper.If(expr == "v2", 2, 0))
```

### Date

```go
// Get the current date and time string
tim, _ := gohelper.NowDateTime("Asia/Shanghai", "2006-01-02 15:04:05")

tim.ToString() -> 2024-09-12 12:02:15
tim.Year() -> 2024
tim.Month() -> September
tim.Day() -> 12
tim.Hour() -> 12
tim.Minute() -> 2
tim.Second() -> 15
```

### Str

####

```go
// Generate string with options (include uppercase, numbers, and symbols)
gohelper.RandStr(16) -> knvmfcmpfqiqcbrh

gohelper.RandStrWithOption(16, Option{
  IncludeNumber:    true,
  IncludeUppercase: true,
  IncludeSymbol:    true,
}) -> nD>fKDvaF\R+1h.G
```
