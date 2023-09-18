package main

import (
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var pasteText string

func main() {
	a := app.New()
	w := a.NewWindow("Autotyper")
	//filePathIcon := widget.NewFileIcon(nil)
	fileLabel := widget.NewLabel("Select file")

	btn := widget.NewButton("[NO FILE]", func() {
		if pasteText != "" {
			fmt.Print(pasteText)

		} else {
			dialog.ShowInformation("БАЛБЕС", "выбери файл", w)
		}
	})
	filePathField := widget.NewButton("Select File", func() {
		dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {
			if err != nil {
				d := dialog.NewError(err, w)
				d.Show()
			}
			if file != nil {
				fileLabel.SetText("Loaded: " + file.URI().String())
				data, _ := io.ReadAll(file)
				pasteText = string(data)
				btn.SetText("[START EXECUTION]")

			}

		}, w)
	})
	w.Resize(fyne.NewSize(200, 320))

	w.SetContent(container.NewVBox(fileLabel, filePathField, btn))
	w.ShowAndRun()

}
