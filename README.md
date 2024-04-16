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

```go
if expr {
  true
} else {
  false
}

gohelper.If(expr, true, false)
```
