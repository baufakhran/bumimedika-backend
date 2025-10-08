package util

import "time"

func GetDatetimeFormatID(t time.Time) string {
	format := "02 January 2006 15:04:05"
	return t.Format(format)
}

func GetDateFormatID(t time.Time) string {
	format := "02 January 2006"
	return t.Format(format)
}
