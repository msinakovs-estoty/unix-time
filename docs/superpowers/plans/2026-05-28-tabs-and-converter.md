# Tabs + Converter Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a two-tab layout — "Main" (existing clock) and "Converter" (Unix ↔ Date bidirectional converter with radio toggle) — to the Fyne timer app.

**Architecture:** Extract pure conversion logic into `convert.go` (testable, no GUI), refactor `main.go` into `makeMainTab()` + `makeConverterTab()` + `main()` using `container.NewAppTabs`.

**Tech Stack:** Go, `fyne.io/fyne/v2`

---

### Task 1: Conversion logic with tests

**Files:**
- Create: `convert.go`
- Create: `convert_test.go`

- [ ] **Step 1: Write failing tests**

Create `C:\Users\User\Documents\timer\convert_test.go`:

```go
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
```

- [ ] **Step 2: Run tests to verify they fail**

```
cd C:\Users\User\Documents\timer
go test ./... -run "TestUnixToDate|TestDateToUnix" -v
```
Expected: FAIL — `unixToDate` and `dateToUnix` undefined.

- [ ] **Step 3: Write the implementation**

Create `C:\Users\User\Documents\timer\convert.go`:

```go
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
```

- [ ] **Step 4: Run tests to verify they pass**

```
go test ./... -run "TestUnixToDate|TestDateToUnix" -v
```
Expected: all 4 tests PASS.

- [ ] **Step 5: Commit**

```
git add convert.go convert_test.go
git commit -m "feat: add unix/date conversion functions with tests"
```

---

### Task 2: Refactor main.go with AppTabs

**Files:**
- Modify: `main.go`

- [ ] **Step 1: Replace main.go**

Replace the entire contents of `C:\Users\User\Documents\timer\main.go` with:

```go
package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func makeMainTab() *container.TabItem {
	now := time.Now()

	unixBind := binding.NewString()
	unixBind.Set(fmt.Sprintf("%d", now.Unix()))
	unixLabel := widget.NewLabelWithData(unixBind)
	unixLabel.Alignment = fyne.TextAlignCenter
	unixLabel.TextStyle = fyne.TextStyle{Bold: true, Monospace: true}

	humanBind := binding.NewString()
	humanBind.Set(now.Format("2006-01-02 15:04:05"))
	humanLabel := widget.NewLabelWithData(humanBind)
	humanLabel.Alignment = fyne.TextAlignCenter
	humanLabel.TextStyle = fyne.TextStyle{Monospace: true}

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			t := time.Now()
			unixBind.Set(fmt.Sprintf("%d", t.Unix()))
			humanBind.Set(t.Format("2006-01-02 15:04:05"))
		}
	}()

	return container.NewTabItem("Main", container.NewPadded(container.NewVBox(unixLabel, humanLabel)))
}

func makeConverterTab() *container.TabItem {
	resultLabel := widget.NewLabel("")
	resultLabel.Alignment = fyne.TextAlignCenter
	resultLabel.TextStyle = fyne.TextStyle{Monospace: true}

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter value...")

	direction := "Unix → Date"

	entry.OnChanged = func(s string) {
		if s == "" {
			resultLabel.SetText("")
			return
		}
		if direction == "Unix → Date" {
			resultLabel.SetText(unixToDate(s))
		} else {
			resultLabel.SetText(dateToUnix(s))
		}
	}

	radio := widget.NewRadioGroup([]string{"Unix → Date", "Date → Unix"}, func(selected string) {
		direction = selected
		entry.SetText("")
		resultLabel.SetText("")
	})
	radio.SetSelected("Unix → Date")
	radio.Horizontal = true

	return container.NewTabItem("Converter", container.NewPadded(container.NewVBox(radio, entry, resultLabel)))
}

func main() {
	a := app.New()
	w := a.NewWindow("Timer")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(300, 180))

	w.SetContent(container.NewAppTabs(
		makeMainTab(),
		makeConverterTab(),
	))
	w.ShowAndRun()
}
```

- [ ] **Step 2: Build to verify it compiles**

```
$env:CGO_ENABLED=1; $env:CC="C:\Users\User\AppData\Local\Microsoft\WinGet\Packages\BrechtSanders.WinLibs.POSIX.UCRT_Microsoft.Winget.Source_8wekyb3d8bbwe\mingw64\bin\gcc.exe"; go build -ldflags "-H windowsgui" -o timer.exe .
```
Expected: `timer.exe` produced with no errors.

- [ ] **Step 3: Run all tests**

```
go test ./... -v
```
Expected: all tests PASS (the conversion tests from Task 1 still pass against the refactored codebase).

- [ ] **Step 4: Commit**

```
git add main.go
git commit -m "feat: add AppTabs with Main and Converter screens"
```
