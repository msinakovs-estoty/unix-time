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

func main() {
	a := app.New()
	w := a.NewWindow("Timer")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(300, 100))

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

	w.SetContent(container.NewPadded(container.NewVBox(unixLabel, humanLabel)))
	w.ShowAndRun()
}
