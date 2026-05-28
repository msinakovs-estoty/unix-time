package main

import (
	_ "embed"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

//go:embed icon.svg
var iconBytes []byte

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

	return container.NewTabItem("Main", container.NewVBox(unixLabel, humanLabel))
}

func makeConverterTab() *container.TabItem {
	resultLabel := widget.NewLabel("—")
	resultLabel.Alignment = fyne.TextAlignCenter
	resultLabel.TextStyle = fyne.TextStyle{Monospace: true}

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter value...")

	direction := "Unix → Date"

	entry.OnChanged = func(s string) {
		if s == "" {
			resultLabel.SetText("—")
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
		resultLabel.SetText("—")
	})
	radio.SetSelected("Unix → Date")
	radio.Horizontal = true

	return container.NewTabItem("Converter", container.NewVBox(radio, entry, resultLabel))
}

func main() {
	a := app.New()
	a.SetIcon(fyne.NewStaticResource("icon.svg", iconBytes))
	w := a.NewWindow("Timer")
	w.SetContent(container.NewAppTabs(
		makeMainTab(),
		makeConverterTab(),
	))
	w.Resize(fyne.NewSize(300, 130))
	w.ShowAndRun()
}
