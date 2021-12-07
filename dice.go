package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showDice() {

	img := canvas.NewImageFromFile("images\\dice1.png")
	img.FillMode = canvas.ImageFill(canvas.ImageScaleFastest)
	btn := widget.NewButton("Play", func() {
		rand := rand.Intn(6) + 1
		img.File = fmt.Sprintf("images\\dice%d.png", rand)
		img.Refresh()
	})
	diceContainer := container.NewVBox(
		img,
		btn,
	)
	w := myApp.NewWindow("Dice Game")
	w.Resize(fyne.NewSize(500, 200))
	w.SetContent(container.NewBorder(DeskBtn, nil, nil, nil, diceContainer))
	w.Resize(fyne.NewSize(300, 300))

	w.Show()
}
