package main

import (
	"testing"
	"time"
)

func TestUnixToDate_invalid(t *testing.T) {
	if got := unixToDate("abc"); got != "invalid input" {
		t.Errorf("got %q, want %q", got, "invalid input")
	}
	if got := unixToDate(""); got != "invalid input" {
		t.Errorf("got %q, want %q", got, "invalid input")
	}
}

func TestUnixToDate_valid(t *testing.T) {
	got := unixToDate("1000000000")
	want := time.Unix(1000000000, 0).Local().Format("2006-01-02 15:04:05")
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestDateToUnix_roundtrip(t *testing.T) {
	date := time.Unix(1000000000, 0).Local().Format("2006-01-02 15:04:05")
	if got := dateToUnix(date); got != "1000000000" {
		t.Errorf("got %q, want %q", got, "1000000000")
	}
}

func TestDateToUnix_invalid(t *testing.T) {
	if got := dateToUnix("not-a-date"); got != "invalid input" {
		t.Errorf("got %q, want %q", got, "invalid input")
	}
	if got := dateToUnix(""); got != "invalid input" {
		t.Errorf("got %q, want %q", got, "invalid input")
	}
}
