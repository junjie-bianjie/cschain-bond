package utils

import (
	"strconv"
	"strings"
	"time"
)

// Gets the first and last day of the month, input example "2020-08"  "2013-01" "1998-3"
func GetMonthStartAndEnd(yearAndMonth string) (string, string) {
	// todo add a regex to limit the input
	split := strings.Split(yearAndMonth, "-")
	year := split[0]
	month := split[1]

	yInt, _ := strconv.Atoi(year)
	yMonth, _ := strconv.Atoi(month)
	timeLayout := "2006-01-02"

	startDate := time.Date(yInt, time.Month(yMonth), 1, 0, 0, 0, 0, time.Local).Format(timeLayout)
	endDate := time.Date(yInt, time.Month(yMonth+1), 0, 0, 0, 0, 0, time.Local).Format(timeLayout)
	return startDate, endDate
}

func String2Time(str string) time.Time {
	timeLayout := "2006-01-02"
	res, _ := time.ParseInLocation(timeLayout, str, time.Local)
	return res
}
