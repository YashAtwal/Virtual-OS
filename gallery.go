package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showGalleryApp(w fyne.Window) {

	root_src := "C:\\Users\\Skillzzzzzy\\Pictures"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}
	tabs := container.NewAppTabs()

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpg" || extension == "jpeg" {
				image := canvas.NewImageFromFile(root_src + "\\" + file.Name())
				//image.FillMode = canvas.ImageFillContain
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}

	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(container.NewBorder(nil, panelContent, nil, nil, tabs))
	w.Resize(fyne.NewSize(1280, 720))
	w.Show()
}
