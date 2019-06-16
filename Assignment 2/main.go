package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Text Editor"),
		widget.NewMultiLineEntry(),
		widget.NewMultiLineEntry(),
		widget.NewMultiLineEntry(),
		widget.NewButton("Exit", func() {
			app.Quit()
		}),
	))
	w.Resize(fyne.NewSize(480, 320))

	w.ShowAndRun()
}
