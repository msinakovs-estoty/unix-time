package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Timer")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(300, 100))

	unixLabel := widget.NewLabelWithStyle(
		fmt.Sprintf("%d", time.Now().Unix()),
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true, Monospace: true},
	)

	humanLabel := widget.NewLabelWithStyle(
		time.Now().Format("2006-01-02 15:04:05"),
		fyne.TextAlignCenter,
		fyne.TextStyle{Monospace: true},
	)

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			now := time.Now()
			unixLabel.SetText(fmt.Sprintf("%d", now.Unix()))
			humanLabel.SetText(now.Format("2006-01-02 15:04:05"))
		}
	}()

	w.SetContent(container.NewPadded(container.NewVBox(unixLabel, humanLabel)))
	w.ShowAndRun()
}
