package pkg

import (
	"strconv"
	"time"
)

var OneHour = time.Duration(1) * time.Hour

func GetTimeOutSeconds(durations ...int) time.Duration {
	if len(durations) == 0 {
		return time.Duration(5) * time.Second
	}
	return time.Duration(durations[0]) * time.Second
}

func GetTimestamp(t *time.Time) string {
	if t == nil {
		return strconv.FormatInt(time.Now().Unix(), 10)
	}
	return strconv.FormatInt(t.Unix(), 10)
}

func GetDay(t *time.Time) string {
	if t == nil {
		return time.Now().Format("20060102")
	}
	return t.Format("20060102")
}
func GetMin(t *time.Time) string {
	if t == nil {
		return time.Now().Format("200601021504")
	}
	return t.Format("200601021504")
}
func GetSec(t *time.Time) string {
	if t == nil {
		return time.Now().Format("2006-01-02T15:04:05")
	}
	return t.Format("2006-01-02T15:04:05")
}
