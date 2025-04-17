package gokit

import (
	"github.com/imajinyun/gokit/internal/datetime"
)

func NewDateTime(date string, zone string, layout string) (*datetime.DateTime, error) {
	return datetime.NewDateTime(date, zone, layout)
}

func NowDateTime(zone string, layout string) (*datetime.DateTime, error) {
	return datetime.NowDateTime(zone, layout)
}
