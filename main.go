package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyneTest/theme"
	"fyneTest/ui"
)

// background color
func main() {
	fmt.Println("main")
	
	a := app.New()
	w := a.NewWindow("DockPanel")
	a.Settings().SetTheme(&theme.MyTheme{})
	mainUi := ui.CreateMainUi(w)
	
	//rect := canvas.NewRectangle(color.Gray{Y: 128})
	w.SetContent(mainUi.WindowContent)
	w.SetPadded(false)
	
	iconRes, _ := fyne.LoadResourceFromPath("go.png")
	w.SetIcon(iconRes)
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	w.SetMaster()
	w.ShowAndRun()
	
}


