package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("PeP OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget

var img fyne.CanvasObject
var DeskBtn fyne.Widget

var panelContent *fyne.Container

func main() {
	myApp.Settings().SetTheme(theme.LightTheme())
	img = canvas.NewImageFromFile("C:\\Users\\Skillzzzzzy\\Pictures\\Wallpapers\\esQ1FjZ.jpg")

	btn1 = widget.NewButtonWithIcon("Weather", theme.InfoIcon(), func() {
		showWeatherApp(myWindow)
	})
	btn2 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		showCalc()
	})
	btn3 = widget.NewButtonWithIcon("Gallery", theme.MediaVideoIcon(), func() {
		showGalleryApp(myWindow)
	})
	btn4 = widget.NewButtonWithIcon("Text Editor", theme.DocumentIcon(), func() {
		showTextEditor()
	})
	btn5 = widget.NewButtonWithIcon("Dice Roll", theme.MediaPlayIcon(), func() {
		showDice()
	})
	btn6 = widget.NewButtonWithIcon("Clock", theme.ComputerIcon(), func() {
		showClock()
	})
	DeskBtn = widget.NewButtonWithIcon("Desktop", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(nil, panelContent, nil, nil, img))
	})
	panelContent = container.NewVBox((container.NewGridWithColumns(7, DeskBtn, btn6, btn1, btn2, btn3, btn4, btn5)))

	myWindow.Resize(fyne.NewSize(1280, 880))
	myWindow.CenterOnScreen()

	myWindow.SetContent(
		container.NewBorder(nil, panelContent, nil, nil, img),
	)
	myWindow.ShowAndRun()
}
