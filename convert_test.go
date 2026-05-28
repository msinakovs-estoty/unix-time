package main

import (
	"fmt"
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
	if got := unixToDate("1000000000"); got == "invalid input" {
		t.Errorf("unixToDate(1000000000) = %q, expected a valid date string", got)
	}
}

func TestDateToUnix_roundtrip(t *testing.T) {
	ts := int64(1000000000)
	date := unixToDate(fmt.Sprintf("%d", ts))
	if date == "invalid input" {
		t.Fatalf("unixToDate(%d) unexpectedly returned invalid input", ts)
	}
	got := dateToUnix(date)
	want := fmt.Sprintf("%d", ts)
	if got != want {
		t.Errorf("round-trip: got %q, want %q", got, want)
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
