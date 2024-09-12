package gohelper

import (
	"fmt"
	"testing"
	"time"
)

func TestNowDateTime_BeginEndOfDay(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now
		outBegin := time.Date(out.Year(), out.Month(), out.Day(), 0, 0, 0, 0, loc).Format(tt.layout)
		outEnd := time.Date(out.Year(), out.Month(), out.Day(), 23, 59, 59, 0, loc).Format(tt.layout)

		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}
		gotBegin, gotEnd := tim.BeginOfDay().Format(tt.layout), tim.EndOfDay().Format(tt.layout)

		if gotBegin != outBegin {
			t.Errorf("tim.BeginEndOfDay() = %s, want: %s", gotBegin, outBegin)
		}

		if gotEnd != outEnd {
			t.Errorf("tim.BeginEndOfDay() = %s, want: %s", gotEnd, outEnd)
		}
	}
}

func TestNowDateTime_BeginEndOfMonth(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		tmp := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		out := time.Date(tmp.Year(), tmp.Month(), 1, 0, 0, 0, 0, tmp.Location())
		outBegin := out.Format(tt.layout)
		outEnd := out.AddDate(0, 1, 0).Add(time.Duration(-1) * time.Second).Format(tt.layout)

		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}
		gotBegin, gotEnd := tim.BeginOfMonth().Format(tt.layout), tim.EndOfMonth().Format(tt.layout)

		if gotBegin != outBegin {
			t.Errorf("tim.BeginEndOfMonth() = %s, want: %s", gotBegin, outBegin)
		}

		if gotEnd != outEnd {
			t.Errorf("tim.BeginEndOfMonth() = %s, want: %s", gotEnd, outEnd)
		}
	}
}

func TestNowDateTime_BeginEndOfYear(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		tmp := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		out := time.Date(tmp.Year(), 1, 1, 0, 0, 0, 0, tmp.Location())
		outBegin := out.Format(tt.layout)
		outEnd := out.AddDate(1, 0, 0).Add(time.Duration(-1) * time.Second).Format(tt.layout)

		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}
		gotBegin, gotEnd := tim.BeginOfYear().Format(tt.layout), tim.EndOfYear().Format(tt.layout)

		if gotBegin != outBegin {
			t.Errorf("tim.BeginEndOfYear() = %s, want: %s", gotBegin, outBegin)
		}

		if gotEnd != outEnd {
			t.Errorf("tim.BeginEndOfYear() = %s, want: %s", gotEnd, outEnd)
		}
	}
}

func TestNowDateTime_Day(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Day()
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.Day(); got != out {
			t.Errorf("tim.Day() = %d, want: %d", got, out)
		}
	}
}

func TestNowDateTime_FmtDate(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Format(time.DateOnly)
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.FmtDate(); got != out {
			t.Errorf("tim.FmtDate() = %s, want: %s", got, out)
		}
	}
}

func TestNowDateTime_FmtDateTime(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Format(time.DateTime)
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.FmtDateTime(); got != out {
			t.Errorf("tim.FmtDateTime() = %s, want: %s", got, out)
		}
	}
}

func TestNowDateTime_FmtTime(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Format(time.TimeOnly)
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.FmtTime(); got != out {
			t.Errorf("tim.FmtTime() = %s, want: %s", got, out)
		}
	}
}

func TestNowDateTime_Hour(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Hour()
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.Hour(); got != out {
			t.Errorf("tim.Hour() = %d, want: %d", got, out)
		}
	}
}

func TestNowDateTime_Minute(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Minute()
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.Minute(); got != out {
			t.Errorf("tim.Minute() = %d, want: %d", got, out)
		}
	}
}

func TestNowDateTime_Month(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Month()
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.Month(); got != out {
			t.Errorf("tim.Month() = %d, want: %d", got, out)
		}
	}
}

func TestNowDateTime_Now(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Format(tt.layout)
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.Now(); got != out {
			t.Errorf("tim.Now() = %s, want: %s", got, out)
		}
	}
}

func TestNowDateTime_Second(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Second()
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.Second(); got != out {
			t.Errorf("tim.Second() = %d, want: %d", got, out)
		}
	}
}

func TestNowDateTime_Today(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Format(time.DateOnly)
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.Today(); got != out {
			t.Errorf("tim.Today() = %s, want: %s", got, out)
		} else {
			t.Log(got)
		}
	}
}

func TestNowDateTime_Tomorrow(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).AddDate(0, 0, 1).Format(tt.layout)
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		got := tim.Tomorrow().Format(tt.layout)
		if got != out {
			t.Errorf("tim.Tomorrow() = %s, want: %s", got, out)
		}
	}
}

func TestNowDateTime_ToString(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Format(tt.layout)
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.ToString(); got != out {
			t.Errorf("tim.ToString() = %s, want: %s", got, out)
		}
	}
}

func TestNowDateTime_Unix(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Unix()
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.Unix(); got != out {
			t.Errorf("tim.Unix() = %d, want: %d", got, out)
		}
	}
}

func TestNowDateTime_Weekday(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Weekday()
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.Weekday(); got != out {
			t.Errorf("tim.Weekday() = %d, want: %d", got, out)
		}
	}
}

func TestNowDateTime_Year(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).Year()
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.Year(); got != out {
			t.Errorf("tim.Year() = %d, want: %d", got, out)
		}
	}
}

func TestNowDateTime_YearDay(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).YearDay()
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		if got := tim.YearDay(); got != out {
			t.Errorf("tim.Year() = %d, want: %d", got, out)
		}
	}
}

func TestNowDateTime_Yesterday(t *testing.T) {
	now, tests := getNowDateTime()
	for _, tt := range tests {
		loc, err := time.LoadLocation(tt.zone)
		if err != nil {
			t.Error(err)
		}

		out := now.In(loc).AddDate(0, 0, -1).Format(tt.layout)
		tim, err := NowDateTime(tt.zone, tt.layout)
		if err != nil {
			t.Error(err)
		}

		got := tim.Yesterday().Format(tt.layout)
		if got != out {
			t.Errorf("tim.Yesterday() = %s, want: %s", got, out)
		}
	}
}

func TestMain(t *testing.T) {
	tim, _ := NowDateTime("Asia/Shanghai", time.DateTime)
	fmt.Println("tim.ToString() ->", tim.ToString())
	fmt.Println("tim.Year() ->", tim.Year())
	fmt.Println("tim.Month() ->", tim.Month())
	fmt.Println("tim.Day() ->", tim.Day())
	fmt.Println("tim.Hour() ->", tim.Hour())
	fmt.Println("tim.Minute() ->", tim.Minute())
	fmt.Println("tim.Second() ->", tim.Second())
	fmt.Println("time.Nanosecond() ->", tim.Nanosecond())
	fmt.Println("tim.Weekday() ->", tim.Weekday())
	fmt.Println("tim.Unix() ->", tim.Unix())
	fmt.Println("tim.Now() ->", tim.Now())
	fmt.Println("tim.Today() ->", tim.Today())
	fmt.Println("tim.Tomorrow() ->", tim.Tomorrow())
	fmt.Println("tim.Yesterday() ->", tim.Yesterday())
	fmt.Println("tim.FmtDate() ->", tim.FmtDate())
	fmt.Println("tim.FmtTime() ->", tim.FmtTime())
	fmt.Println("tim.FmtDateTime() ->", tim.FmtDateTime())
	fmt.Println("tim.BeginOfDay() ->", tim.BeginOfDay())
	fmt.Println("tim.EndOfDay() ->", tim.EndOfDay())
}

func getNowDateTime() (time.Time, []struct {
	in     time.Time
	zone   string
	layout string
}) {
	now := time.Now()
	var tests = []struct {
		in     time.Time
		zone   string
		layout string
	}{
		{now, "Asia/Shanghai", time.DateTime},
		{now, "Asia/Singapore", time.RFC3339},
		{now, "America/New_York", time.RFC1123},
	}

	return now, tests
}
