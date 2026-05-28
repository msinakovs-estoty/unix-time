# Tabs + Converter Design

## Summary
Add a two-tab layout to the timer app. Tab 1 ("Main") keeps the existing live clock. Tab 2 ("Converter") lets the user convert between Unix timestamp and human-readable date/time in either direction.

## Navigation
`container.NewAppTabs` with two tabs at the top. Window resizes to 300×180 to accommodate the extra UI. Fixed size remains on.

## Main Tab
Unchanged: two live-updating labels (Unix timestamp bold+monospace, human-readable monospace), centered, refreshing every second via data bindings.

## Converter Tab
- `widget.RadioGroup` — horizontal, options: `"Unix → Date"` and `"Date → Unix"`. Default: `"Unix → Date"`.
- `widget.Entry` — single-line text input for the value to convert.
- `widget.Label` — read-only result, updated on every keystroke via `Entry.OnChanged`.
- Conversion logic:
  - `"Unix → Date"`: parse input as `int64`, call `time.Unix(n, 0).Format("2006-01-02 15:04:05")`.
  - `"Date → Unix"`: parse input with `time.Parse("2006-01-02 15:04:05", s)`, output `.Unix()` as string.
  - Invalid input → result label shows `"invalid input"`.
- Switching the radio clears the entry and result label.

## Architecture
Single `main.go` split into focused functions:
- `makeMainTab() *container.TabItem` — builds the Main tab, starts the ticker goroutine.
- `makeConverterTab() *container.TabItem` — builds the Converter tab, pure UI/logic, no goroutines.
- `main()` — creates app, window, assembles tabs, calls `ShowAndRun`.

The ticker goroutine from Main is started inside `makeMainTab` and runs for the lifetime of the app (no explicit cancel needed for a single-window app).

## File Changes
- Modify: `main.go` (only file)
