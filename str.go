package gokit

import (
	"github.com/imajinyun/gokit/internal/strs"
)

func RandStr(n int) string {
	return strs.RandStr(n)
}

func RandStrWithOption(n int, opt strs.Option) string {
	return strs.RandStrWithOption(n, opt)
}

func Empty(val any) bool {
	return strs.Empty(val)
}

func Trim(s string, char string) string {
	return strs.Trim(s, char)
}

func TrimLeft(s string, char string) string {
	return strs.TrimLeft(s, char)
}

func TrimRight(s string, char string) string {
	return strs.TrimRight(s, char)
}
