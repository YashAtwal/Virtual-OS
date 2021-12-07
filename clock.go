package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showTime(clock *widget.Label) {
	formatted := time.Now().Format("3:04:05")
	clock.SetText(formatted)
}
func showClock() {
	clock := widget.NewLabel("Current Time")
	showTime(clock)
	img := canvas.NewImageFromFile("images\\time.png")
	img.FillMode = canvas.ImageFill(canvas.ImageScaleFastest)
	clockContainer := container.NewVBox(
		img,
		clock,
	)
	go func() {
		t := time.NewTicker(time.Second)
		for range t.C {
			showTime(clock)
		}
	}()
	w := myApp.NewWindow("Ghadiii")
	w.SetContent(container.NewBorder(DeskBtn, nil, nil, nil, clockContainer))
	w.Resize(fyne.NewSize(200, 200))
	w.Show()

}
