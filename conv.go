package gokit

import "github.com/imajinyun/gohelper/internal/conv"

func ToString(data any) string {
	return conv.ToString(data)
}

func ToJson(data any) (string, error) {
	return conv.ToJson(data)
}

func StructToMap[T any](data any) map[string]T {
	return conv.StructToMap[T](data)
}
