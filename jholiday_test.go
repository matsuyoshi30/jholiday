package jholiday_test

import (
	"testing"
	"time"

	"github.com/matsuyoshi30/jholiday"
)

func TestHolidayName(t *testing.T) {
	jst := jholiday.GetJST()

	tests := []struct {
		date time.Time
		want string
	}{
		{
			date: time.Date(2024, time.January, 1, 0, 0, 0, 0, jst),
			want: "元旦",
		},
		{
			date: time.Date(2024, time.January, 8, 0, 0, 0, 0, jst),
			want: "成人の日",
		},
		{
			date: time.Date(1999, time.January, 15, 0, 0, 0, 0, jst),
			want: "成人の日",
		},
		{
			date: time.Date(2024, time.January, 15, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(2024, time.February, 11, 0, 0, 0, 0, jst),
			want: "建国記念の日",
		},
		{
			date: time.Date(2024, time.February, 23, 0, 0, 0, 0, jst),
			want: "天皇誕生日",
		},
		{
			date: time.Date(1989, time.February, 24, 0, 0, 0, 0, jst),
			want: "昭和天皇の大喪の礼",
		},
		{
			date: time.Date(2024, time.February, 24, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(2024, time.March, 20, 0, 0, 0, 0, jst),
			want: "春分の日",
		},
		{
			date: time.Date(2024, time.March, 21, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(2024, time.April, 29, 0, 0, 0, 0, jst),
			want: "昭和の日",
		},
		{
			date: time.Date(1989, time.April, 29, 0, 0, 0, 0, jst),
			want: "みどりの日",
		},
		{
			date: time.Date(1988, time.April, 29, 0, 0, 0, 0, jst),
			want: "天皇誕生日",
		},
		{
			date: time.Date(2019, time.April, 30, 0, 0, 0, 0, jst),
			want: "国民の休日",
		},
		{
			date: time.Date(1959, time.April, 10, 0, 0, 0, 0, jst),
			want: "皇太子明仁親王の結婚の儀",
		},
		{
			date: time.Date(2024, time.April, 10, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(2024, time.May, 3, 0, 0, 0, 0, jst),
			want: "憲法記念日",
		},
		{
			date: time.Date(2024, time.May, 4, 0, 0, 0, 0, jst),
			want: "みどりの日",
		},
		{
			date: time.Date(2006, time.May, 4, 0, 0, 0, 0, jst),
			want: "国民の休日",
		},
		{
			date: time.Date(2024, time.May, 5, 0, 0, 0, 0, jst),
			want: "こどもの日",
		},
		{
			date: time.Date(2024, time.May, 6, 0, 0, 0, 0, jst),
			want: "振替休日",
		},
		{
			date: time.Date(2019, time.May, 1, 0, 0, 0, 0, jst),
			want: "即位の日",
		},
		{
			date: time.Date(2019, time.May, 2, 0, 0, 0, 0, jst),
			want: "国民の休日",
		},
		{
			date: time.Date(2024, time.May, 2, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(1993, time.June, 9, 0, 0, 0, 0, jst),
			want: "皇太子徳仁親王の結婚の儀",
		},
		{
			date: time.Date(2024, time.June, 9, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(2024, time.July, 15, 0, 0, 0, 0, jst),
			want: "海の日",
		},
		{
			date: time.Date(2021, time.July, 22, 0, 0, 0, 0, jst),
			want: "海の日",
		},
		{
			date: time.Date(2021, time.July, 23, 0, 0, 0, 0, jst),
			want: "スポーツの日",
		},
		{
			date: time.Date(2020, time.July, 23, 0, 0, 0, 0, jst),
			want: "海の日",
		},
		{
			date: time.Date(2020, time.July, 24, 0, 0, 0, 0, jst),
			want: "スポーツの日",
		},
		{
			date: time.Date(2019, time.July, 15, 0, 0, 0, 0, jst),
			want: "海の日",
		},
		{
			date: time.Date(1996, time.July, 20, 0, 0, 0, 0, jst),
			want: "海の日",
		},
		{
			date: time.Date(2024, time.July, 20, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(2024, time.August, 11, 0, 0, 0, 0, jst),
			want: "山の日",
		},
		{
			date: time.Date(2021, time.August, 8, 0, 0, 0, 0, jst),
			want: "山の日",
		},
		{
			date: time.Date(2020, time.August, 10, 0, 0, 0, 0, jst),
			want: "山の日",
		},
		{
			date: time.Date(2019, time.August, 11, 0, 0, 0, 0, jst),
			want: "山の日",
		},
		{
			date: time.Date(2024, time.August, 10, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(2024, time.September, 22, 0, 0, 0, 0, jst),
			want: "秋分の日",
		},
		{
			date: time.Date(2024, time.September, 23, 0, 0, 0, 0, jst),
			want: "振替休日",
		},
		{
			date: time.Date(2015, time.September, 22, 0, 0, 0, 0, jst),
			want: "国民の休日",
		},
		{
			date: time.Date(1966, time.September, 15, 0, 0, 0, 0, jst),
			want: "敬老の日",
		},
		{
			date: time.Date(2024, time.September, 15, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(2024, time.October, 14, 0, 0, 0, 0, jst),
			want: "スポーツの日",
		},
		{
			date: time.Date(2021, time.October, 14, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(2020, time.October, 14, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(2019, time.October, 14, 0, 0, 0, 0, jst),
			want: "体育の日",
		},
		{
			date: time.Date(2019, time.October, 22, 0, 0, 0, 0, jst),
			want: "即位礼正殿の儀",
		},
		{
			date: time.Date(1967, time.October, 10, 0, 0, 0, 0, jst),
			want: "体育の日",
		},
		{
			date: time.Date(2024, time.October, 10, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(2024, time.November, 3, 0, 0, 0, 0, jst),
			want: "文化の日",
		},
		{
			date: time.Date(2024, time.November, 23, 0, 0, 0, 0, jst),
			want: "勤労感謝の日",
		},
		{
			date: time.Date(1990, time.November, 12, 0, 0, 0, 0, jst),
			want: "即位礼正殿の儀",
		},
		{
			date: time.Date(2024, time.November, 12, 0, 0, 0, 0, jst),
			want: "",
		},
		{
			date: time.Date(2018, time.December, 23, 0, 0, 0, 0, jst),
			want: "天皇誕生日",
		},
		{
			date: time.Date(2024, time.December, 23, 0, 0, 0, 0, jst),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := jholiday.HolidayName(tt.date)
			if tt.want != got {
				t.Fatalf("%v: want %s but got %s", tt.date, tt.want, got)
			}
		})
	}
}
