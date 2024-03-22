package util

import "time"

const layout = "2006-01-02T15:04:05"

func StringToTime(string) time.Time {
	convertedeTime, _ := time.Parse(layout, "")
	return convertedeTime
}
