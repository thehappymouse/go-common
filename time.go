package utils

import (
	"time"
)

const (
	DataTimeMilli  = "2006-01-02T15:04:05.000"
	DataTimeMilli2 = "2006-01-02T150405.000"
	DateTimeDate   = "20060102"
)

// 工作日计算
func WorkDayAdd(days int, startDay time.Time) time.Time {
	for i := 0; i < days; i++ {
		startDay = startDay.AddDate(0, 0, 1)
		if startDay.Weekday() == time.Sunday || startDay.Weekday() == time.Saturday {
			i--
		}
	}
	return startDay
}
