package jholiday

// https://addinbox.sakura.ne.jp/holiday_logic.htm

import (
	"math"
	"time"
)

var jst *time.Location

func init() {
	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
}

func IsHoliday(date time.Time) bool {
	return len(HolidayName(date)) > 0
}

func HolidayName(date time.Time) string {
	var (
		EnforcementSubstituteHoliday = time.Date(1973, time.April, 12, 0, 0, 0, 0, jst)
	)

	date = date.In(jst)

	year := date.Year()
	month := date.Month()
	day := date.Day()

	date = time.Date(year, month, day, 0, 0, 0, 0, jst)
	holidayName := holiday(date)

	if len(holidayName) == 0 {
		if date.Weekday() == time.Monday {
			if date.After(EnforcementSubstituteHoliday) {
				holidayName := holiday(date.Add(-1))
				if len(holidayName) > 0 {
					return "振替休日"
				}
			}
		}
	}

	return holidayName
}

func holiday(date time.Time) string {
	var (
		EnforcementPublicHolidayLaw        = time.Date(1948, time.July, 20, 0, 0, 0, 0, jst)     // 祝日法施行
		ShowaEmperorGreatMourningReception = time.Date(1989, time.February, 24, 0, 0, 0, 0, jst) // 昭和天皇の大喪の礼
		PrinceAkihitoWeddingCeremony       = time.Date(1959, time.April, 10, 0, 0, 0, 0, jst)    // 明仁親王の結婚の儀
		PrinceTokujinWeddingCeremony       = time.Date(1993, time.June, 9, 0, 0, 0, 0, jst)      // 徳仁親王の結婚の儀
		EnthronementCeremony               = time.Date(1990, time.November, 12, 0, 0, 0, 0, jst) // 即位礼正殿の儀
		AbdicationEmperorHeisei            = time.Date(2019, time.April, 30, 0, 0, 0, 0, jst)    // 平成天皇の退位
		PrinceTokujinAccessionThrone       = time.Date(2019, time.May, 1, 0, 0, 0, 0, jst)       // 徳仁親王の即位
		GW2019NationalHoliday              = time.Date(2019, time.May, 2, 0, 0, 0, 0, jst)       // 2019GW国民の休日
		PrinceTokujinAccessionCeremoty     = time.Date(2019, time.October, 22, 0, 0, 0, 0, jst)  // 即位礼正殿の儀_徳仁親王
	)

	if date.Before(EnforcementPublicHolidayLaw) {
		return ""
	}

	year := date.Year()
	month := date.Month()
	day := date.Day()

	switch month {
	case time.January:
		if day == 1 {
			return "元旦"
		}
		if year >= 2000 {
			if ((day-1)/7+1) == 2 && date.Weekday() == time.Monday {
				return "成人の日"
			}
		} else {
			if day == 15 {
				return "成人の日"
			}
		}
	case time.February:
		if day == 11 {
			if year >= 1967 {
				return "建国記念の日"
			}
		}
		if day == 23 {
			if year >= 2020 {
				return "天皇誕生日"
			}
		}
		if date.Equal(ShowaEmperorGreatMourningReception) {
			return "昭和天皇の大喪の礼"
		}
	case time.March:
		if day == vernalEquinox(year) {
			return "春分の日"
		}
	case time.April:
		if day == 29 {
			if year >= 2007 {
				return "昭和の日"
			}
			if year >= 1989 {
				return "みどりの日"
			}
			return "天皇誕生日"
		}
		if date.Equal(AbdicationEmperorHeisei) {
			return "国民の休日"
		}
		if date.Equal(PrinceAkihitoWeddingCeremony) {
			return "皇太子明仁親王の結婚の儀"
		}
	case time.May:
		if day == 3 {
			return "憲法記念日"
		}
		if day == 4 {
			if year >= 2007 {
				return "みどりの日"
			}
			if year >= 1986 {
				if date.Weekday() > time.Monday {
					return "国民の休日"
				}
			}
		}
		if day == 5 {
			return "こどもの日"
		}
		if day == 6 {
			if year >= 2007 {
				wd := date.Weekday()
				if wd == time.Tuesday || wd == time.Wednesday {
					return "振替休日"
				}
			}
		}
		if year == 2019 {
			if date.Equal(PrinceTokujinAccessionThrone) {
				return "即位の日"
			}
			if date.Equal(GW2019NationalHoliday) {
				return "国民の休日"
			}
		}
	case time.June:
		if date.Equal(PrinceTokujinWeddingCeremony) {
			return "皇太子徳仁親王の結婚の儀"
		}
	case time.July:
		if year >= 2022 {
			if (((day-1)/7)+1) == 3 && date.Weekday() == time.Monday {
				return "海の日"
			}
		} else if year == 2021 {
			if day == 22 {
				return "海の日"
			}
			if day == 23 {
				return "スポーツの日"
			}
		} else if year == 2020 {
			if day == 23 {
				return "海の日"
			}
			if day == 24 {
				return "スポーツの日"
			}
		} else if year >= 2003 {
			if (((day-1)/7)+1) == 3 && date.Weekday() == time.Monday {
				return "海の日"
			}
		} else if year >= 1996 {
			if day == 20 {
				return "海の日"
			}
		}
	case time.August:
		if year >= 2022 {
			if day == 11 {
				return "山の日"
			}
		} else if year == 2021 {
			if day == 8 {
				return "山の日"
			}
		} else if year == 2020 {
			if day == 10 {
				return "山の日"
			}
		} else if year >= 2016 {
			if day == 11 {
				return "山の日"
			}
		}
	case time.September:
		ad := autumnalEquinox(year)
		if day == ad {
			return "秋分の日"
		}
		if year >= 2003 {
			if (((day-1)/7)+1) == 3 && date.Weekday() == time.Monday {
				return "敬老の日"
			}
			if date.Weekday() == time.Tuesday {
				if day == ad-1 {
					return "国民の休日"
				}
			}
		} else if year >= 1966 {
			if day == 15 {
				return "敬老の日"
			}
		}
	case time.October:
		if year >= 2022 {
			if (((day-1)/7)+1) == 2 && date.Weekday() == time.Monday {
				return "スポーツの日"
			}
		} else if year == 2021 || year == 2020 {
			// 2021年はオリンピック特措法により「スポーツの日」が 7/23 に移動
			// 2020年はオリンピック特措法により「スポーツの日」が 7/24 に移動
		} else if year >= 2000 {
			if (((day-1)/7)+1) == 2 && date.Weekday() == time.Monday {
				return "体育の日"
			}
			if date.Equal(PrinceTokujinAccessionCeremoty) {
				return "即位礼正殿の儀"
			}
		} else if year >= 1966 {
			if day == 10 {
				return "体育の日"
			}
		}
	case time.November:
		if day == 3 {
			return "文化の日"
		}
		if day == 23 {
			return "勤労感謝の日"
		}
		if date.Equal(EnthronementCeremony) {
			return "即位礼正殿の儀"
		}
	case time.December:
		if day == 23 {
			if year >= 1989 && year <= 2018 {
				return "天皇誕生日"
			}
		}
	}

	return ""
}

func vernalEquinox(year int) int {
	if year <= 1947 {
		return 99
	}

	yearf := float64(year)
	if year <= 1979 {
		return int(fix(20.8357 + (0.242194 * (yearf - 1980)) - fix((yearf-1983)/4)))
	}
	if year <= 2099 {
		return int(fix(20.8431 + (0.242194 * (yearf - 1980)) - fix((yearf-1980)/4)))
	}
	if year <= 2150 {
		return int(fix(21.851 + (0.242194 * (yearf - 1980)) - fix((yearf-1980)/4)))
	}
	return 99
}

func autumnalEquinox(year int) int {
	if year <= 1947 {
		return 99
	}

	yearf := float64(year)
	if year <= 1979 {
		return int(fix(23.2588 + (0.242194 * (yearf - 1980)) - fix((yearf-1983)/4)))
	}
	if year <= 2099 {
		return int(fix(23.2488 + (0.242194 * (yearf - 1980)) - fix((yearf-1980)/4)))
	}
	if year <= 2150 {
		return int(fix(24.2488 + (0.242194 * (yearf - 1980)) - fix((yearf-1980)/4)))
	}
	return 99
}

func fix(x float64) float64 {
	if x > 0 {
		return math.Floor(x)
	}
	return math.Ceil(x)
}
