package main

import (
	"testing"
	"time"
)

func TestGetFebruaryDay(t *testing.T) {
	leapYear := 2000
	averageYear := 2022
	if getFebruaryDay(leapYear) != febLeapYearDay {
		t.Fatal("feb leap year day error")
	}
	if getFebruaryDay(averageYear) != febAverageYearDay {
		t.Fatal("feb leap year day error")
	}
}

func TestGetDaysOfMonth(t *testing.T) {
	days := getDaysOfMonth(2022, time.April)
	if days != month30Day {
		t.Fatal("getDaysOfMonth error", days)
	}
	t.Logf("2022年4月天数是:%d\n", days)
	days = getDaysOfMonth(2022, time.May)
	if days != month31Day {
		t.Fatal("getDayOfMonth error", days)
	}
	t.Logf("2022年5月天数是:%d\n", days)
	days = getDaysOfMonth(2000, time.February)
	if days != febLeapYearDay {
		t.Fatal("getDayOfMonth error", days)
	}
	t.Logf("2000年2月天数是:%d\n", days)
	days = getDaysOfMonth(2022, time.February)
	if days != febAverageYearDay {
		t.Fatal("getDayOfMonth error", days)
	}
	t.Logf("2022年2月天数是:%d\n", days)
}
