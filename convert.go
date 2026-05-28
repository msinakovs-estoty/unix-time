package main

import (
	"fmt"
	"strconv"
	"time"
)

func unixToDate(s string) string {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "invalid input"
	}
	return time.Unix(n, 0).Format("2006-01-02 15:04:05")
}

func dateToUnix(s string) string {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if err != nil {
		return "invalid input"
	}
	return fmt.Sprintf("%d", t.Unix())
}
