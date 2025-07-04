package helper

import "time"

func DateFormatter(date time.Time) string {
	format := date.Format("2006-01-02 15:04:05")
	return format
}
