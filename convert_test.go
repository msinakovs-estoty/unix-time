package main

import (
	"testing"
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
	if got := unixToDate("1000000000"); got != "2001-09-09 01:46:40" {
		t.Errorf("got %q, want %q", got, "2001-09-09 01:46:40")
	}
}

func TestDateToUnix_roundtrip(t *testing.T) {
	if got := dateToUnix("2001-09-09 01:46:40"); got != "1000000000" {
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
