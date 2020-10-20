package models

import "time"

func unixToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 12:22:22")
}

func Hello(in string) (out string) {
	out = in + "world"
	return
}
