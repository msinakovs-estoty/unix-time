# Timer GUI Design

## Summary
A small native desktop window (Fyne) that displays the current Unix timestamp and human-readable date/time, updating every second.

## UI
- Window size: ~300×120px, always-on-top
- Label 1 (large, bold): Unix timestamp — e.g. `1748428800`
- Label 2 (smaller): Human-readable — e.g. `2026-05-28 12:00:00`
- No controls, just a display

## Architecture
- Single `main.go` file
- Fyne window with two `widget.Label` elements
- Goroutine-free: use `fyne.io` canvas refresh on a ticker or just update labels in a `go` routine every second

## Dependencies
- `fyne.io/fyne/v2`
