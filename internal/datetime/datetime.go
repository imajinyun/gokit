package datetime

import (
	"time"
)

type DateTime struct {
	Tim time.Time
	Loc *time.Location
	Fmt string
}

func NewDateTime(date string, zone string, layout string) (*DateTime, error) {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return nil, err
	}

	pst, err := time.Parse(layout, date)
	if err != nil {
		return nil, err
	}

	return &DateTime{Tim: pst, Loc: loc, Fmt: layout}, nil
}

func NowDateTime(zone string, layout string) (*DateTime, error) {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return nil, err
	}

	return &DateTime{Tim: time.Now(), Loc: loc, Fmt: layout}, nil
}

func (d *DateTime) Now() string {
	return d.Tim.In(d.Loc).Format(d.Fmt)
}

func (d *DateTime) Today() string {
	return d.Tim.In(d.Loc).Format(time.DateOnly)
}

func (d *DateTime) ToString() string {
	return d.Tim.In(d.Loc).Format(d.Fmt)
}

func (d *DateTime) FmtDate() string {
	return d.Tim.In(d.Loc).Format(time.DateOnly)
}

func (d *DateTime) FmtTime() string {
	return d.Tim.In(d.Loc).Format(time.TimeOnly)
}

func (d *DateTime) FmtDateTime() string {
	return d.Tim.In(d.Loc).Format(time.DateTime)
}

func (d *DateTime) Year() int {
	return d.Tim.In(d.Loc).Year()
}

func (d *DateTime) YearDay() int {
	return d.Tim.In(d.Loc).YearDay()
}

func (d *DateTime) Month() time.Month {
	return d.Tim.In(d.Loc).Month()
}

func (d *DateTime) Day() int {
	return d.Tim.In(d.Loc).Day()
}

func (d *DateTime) Hour() int {
	return d.Tim.In(d.Loc).Hour()
}

func (d *DateTime) Minute() int {
	return d.Tim.In(d.Loc).Minute()
}

func (d *DateTime) Second() int {
	return d.Tim.In(d.Loc).Second()
}

func (d *DateTime) Nanosecond() int {
	return d.Tim.In(d.Loc).Nanosecond()
}

func (d *DateTime) Weekday() time.Weekday {
	return d.Tim.In(d.Loc).Weekday()
}

func (d *DateTime) Unix() int64 {
	return d.Tim.Unix()
}

func (d *DateTime) Tomorrow() time.Time {
	return d.Tim.AddDate(0, 0, 1).In(d.Loc)
}

func (d *DateTime) Yesterday() time.Time {
	return d.Tim.AddDate(0, 0, -1).In(d.Loc)
}
func (d *DateTime) LastWeek() time.Time {
	return d.Tim.AddDate(0, 0, -7).In(d.Loc)
}

func (d *DateTime) NextWeek() time.Time {
	return d.Tim.AddDate(0, 0, 7).In(d.Loc)
}

func (d *DateTime) BeginOfDay() time.Time {
	return time.Date(d.Tim.Year(), d.Tim.Month(), d.Tim.Day(), 0, 0, 0, 0, d.Loc)
}

func (d *DateTime) EndOfDay() time.Time {
	return time.Date(d.Tim.Year(), d.Tim.Month(), d.Tim.Day(), 23, 59, 59, 0, d.Loc)
}

func (d *DateTime) BeginOfMonth() time.Time {
	return time.Date(d.Tim.Year(), d.Tim.Month(), 1, 0, 0, 0, 0, d.Loc)
}

func (d *DateTime) EndOfMonth() time.Time {
	last := time.Date(d.Tim.Year(), d.Tim.Month()+1, 1, 0, 0, 0, 0, d.Loc).Add(time.Duration(-1) * time.Second)
	return time.Date(last.Year(), last.Month(), last.Day(), 23, 59, 59, 0, d.Loc)
}

func (d *DateTime) BeginOfYear() time.Time {
	return time.Date(d.Tim.Year(), 1, 1, 0, 0, 0, 0, d.Loc)
}

func (d *DateTime) EndOfYear() time.Time {
	last := time.Date(d.Tim.Year()+1, 1, 1, 0, 0, 0, 0, d.Loc).Add(time.Duration(-1) * time.Second)
	return time.Date(last.Year(), last.Month(), last.Day(), 23, 59, 59, 0, d.Loc)
}
