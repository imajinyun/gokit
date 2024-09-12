package gohelper

import "time"

type DateTime struct {
	tim time.Time
	loc *time.Location

	layout string
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

	return &DateTime{tim: pst, loc: loc, layout: layout}, nil
}

func NowDateTime(zone string, layout string) (*DateTime, error) {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return nil, err
	}

	return &DateTime{tim: time.Now(), loc: loc, layout: layout}, nil
}

func (d *DateTime) Now() string {
	return d.tim.In(d.loc).Format(d.layout)
}

func (d *DateTime) Today() string {
	return d.tim.In(d.loc).Format(time.DateOnly)
}

func (d *DateTime) ToString() string {
	return d.tim.In(d.loc).Format(d.layout)
}

func (d *DateTime) FmtDate() string {
	return d.tim.In(d.loc).Format(time.DateOnly)
}

func (d *DateTime) FmtTime() string {
	return d.tim.In(d.loc).Format(time.TimeOnly)
}

func (d *DateTime) FmtDateTime() string {
	return d.tim.In(d.loc).Format(time.DateTime)
}

func (d *DateTime) Year() int {
	return d.tim.In(d.loc).Year()
}

func (d *DateTime) YearDay() int {
	return d.tim.In(d.loc).YearDay()
}

func (d *DateTime) Month() time.Month {
	return d.tim.In(d.loc).Month()
}

func (d *DateTime) Day() int {
	return d.tim.In(d.loc).Day()
}

func (d *DateTime) Hour() int {
	return d.tim.In(d.loc).Hour()
}

func (d *DateTime) Minute() int {
	return d.tim.In(d.loc).Minute()
}

func (d *DateTime) Second() int {
	return d.tim.In(d.loc).Second()
}

func (d *DateTime) Nanosecond() int {
	return d.tim.In(d.loc).Nanosecond()
}

func (d *DateTime) Weekday() time.Weekday {
	return d.tim.In(d.loc).Weekday()
}

func (d *DateTime) Unix() int64 {
	return d.tim.Unix()
}

func (d *DateTime) Tomorrow() time.Time {
	return d.tim.AddDate(0, 0, 1).In(d.loc)
}

func (d *DateTime) Yesterday() time.Time {
	return d.tim.AddDate(0, 0, -1).In(d.loc)
}
func (d *DateTime) LastWeek() time.Time {
	return d.tim.AddDate(0, 0, -7).In(d.loc)
}

func (d *DateTime) NextWeek() time.Time {
	return d.tim.AddDate(0, 0, 7).In(d.loc)
}

func (d *DateTime) BeginOfDay() time.Time {
	return time.Date(d.tim.Year(), d.tim.Month(), d.tim.Day(), 0, 0, 0, 0, d.loc)
}

func (d *DateTime) EndOfDay() time.Time {
	return time.Date(d.tim.Year(), d.tim.Month(), d.tim.Day(), 23, 59, 59, 0, d.loc)
}

func (d *DateTime) BeginOfMonth() time.Time {
	return time.Date(d.tim.Year(), d.tim.Month(), 1, 0, 0, 0, 0, d.loc)
}

func (d *DateTime) EndOfMonth() time.Time {
	last := time.Date(d.tim.Year(), d.tim.Month()+1, 1, 0, 0, 0, 0, d.loc).Add(time.Duration(-1) * time.Second)
	return time.Date(last.Year(), last.Month(), last.Day(), 23, 59, 59, 0, d.loc)
}

func (d *DateTime) BeginOfYear() time.Time {
	return time.Date(d.tim.Year(), 1, 1, 0, 0, 0, 0, d.loc)
}

func (d *DateTime) EndOfYear() time.Time {
	last := time.Date(d.tim.Year()+1, 1, 1, 0, 0, 0, 0, d.loc).Add(time.Duration(-1) * time.Second)
	return time.Date(last.Year(), last.Month(), last.Day(), 23, 59, 59, 0, d.loc)
}
