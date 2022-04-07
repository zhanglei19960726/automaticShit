package automaticshit

import "time"

var (
	// day31 有31天的月份
	day31 = map[time.Month]bool{
		time.January:  true,
		time.March:    true,
		time.May:      true,
		time.July:     true,
		time.August:   true,
		time.October:  true,
		time.December: true,
	}
	// day30 有三十天的月份
	day30 = map[time.Month]bool{
		time.April:     true,
		time.June:      true,
		time.September: true,
		time.November:  true,
	}
)

const (
	// febLeapYearDay 二月闰年天数
	febLeapYearDay = 29
	// febAverageYearDay 二月平年的天数
	febAverageYearDay = 28
	// month30Day 30天
	month30Day = 30
	// month31Day 31 天
	month31Day = 31
)

// getDaysOfMonth 获取某个月的天数
func getDaysOfMonth(year int, month time.Month) int {
	if day31[month] {
		return month31Day
	}
	if day30[month] {
		return month30Day
	}
	return getFebruaryDay(year)
}

// getFebruaryDay 获取二月的天数
func getFebruaryDay(year int) int {
	if judgeLeapYears(year) {
		return febLeapYearDay
	}
	return febAverageYearDay
}

// judgeLeapYears 判断是否是闰年
func judgeLeapYears(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
