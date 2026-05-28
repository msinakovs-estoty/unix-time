package main

import (
	"strconv"
	"time"
)

func unixToDate(s string) string {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "invalid input"
	}
	return time.Unix(n, 0).UTC().Format("2006-01-02 15:04:05")
}

func dateToUnix(s string) string {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return "invalid input"
	}
	return strconv.FormatInt(t.Unix(), 10)
}
