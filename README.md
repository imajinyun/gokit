# gohelper

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

### Simplify the judgment of if else to a single line of code:

```go
flag, expr := 0, "some version"
if expr == "v1" {
  flag = 1
} else if expr == "v2" {
  flag = 2
}

flag = gohelper.If(expr == "v1", 1, gohelper.If(expr == "v2", 2, 0))
```

### Generate string with options (include uppercase, numbers, and symbols)

```go
// knvmfcmpfqiqcbrh
gohelper.RandStr(16)

// nD>fKDvaF\R}1h}G
gohelper.RandStrWithOption(16, Option{
  IncludeNumber:    true,
  IncludeUppercase: true,
  IncludeSymbol:    true,
})
```
