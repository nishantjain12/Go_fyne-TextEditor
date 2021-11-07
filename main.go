package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var count int = 0

type tabs struct {
	fileName string
	text     *widget.Entry
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Text Editor")

	tabsNodes := []tabs{}

	Tabs := container.NewDocTabs()

	content := container.New(
		layout.NewVBoxLayout(),
	)

	content.Add(widget.NewButton("Add File", func() {
		count++
		content.Add(
			container.New(
				layout.NewHBoxLayout(),
				widget.NewLabel("-New File"+strconv.Itoa(count)+".txt"),
			),
		)
		textArea := widget.NewMultiLineEntry()
		textArea.SetPlaceHolder("Enter text...")
		entry := tabs{
			fileName: "New File" + strconv.Itoa(count) + ".txt",
			text:     textArea,
		}
		tabsNodes = append(tabsNodes, entry)
		Tabs.Append(container.NewTabItem(entry.fileName, entry.text))
	}))

	hsplitContainer := container.NewHSplit(
		container.NewVBox(
			canvas.NewText("History", color.White),
			content,
		),
		Tabs,
	)
	hsplitContainer.SetOffset(0.2)
	myWindow.SetContent(hsplitContainer)

	fileItem1 := fyne.NewMenuItem("New File", func() {
		count++
		content.Add(
			container.New(
				layout.NewHBoxLayout(),
				widget.NewLabel("-New File"+strconv.Itoa(count)+".txt"),
			),
		)
		textArea := widget.NewMultiLineEntry()
		textArea.SetPlaceHolder("Enter text...")
		entry := tabs{
			fileName: "New File" + strconv.Itoa(count) + ".txt",
			text:     textArea,
		}
		tabsNodes = append(tabsNodes, entry)
		Tabs.Append(container.NewTabItem(entry.fileName, entry.text))
	})

	fileItem2 := fyne.NewMenuItem("Open File", func() {
		openDialog := dialog.NewFileOpen(
			func(uc fyne.URIReadCloser, _e error) {
				body, err := ioutil.ReadAll(uc)
				if err != nil {
					fmt.Print(err)
				} else {
					output := fyne.NewStaticResource("NewFile", body)
					viewTextField := widget.NewMultiLineEntry()

					viewTextField.SetText(string(output.StaticContent))

					entry := tabs{
						fileName: "New File" + strconv.Itoa(count) + ".txt",
						text:     viewTextField,
					}

					tabsNodes = append(tabsNodes, entry)
					Tabs.Append(container.NewTabItem(entry.fileName, entry.text))
					content.Add(container.New(
						layout.NewHBoxLayout(),
						widget.NewLabel("-"+entry.fileName),
					))
				}
			}, myWindow)
		openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openDialog.Show()
	})
	fileItem3 := fyne.NewMenuItem("Save File", func() {
		saveDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _e error) {
				textData := []byte(tabsNodes[Tabs.SelectedIndex()].text.Text)
				uc.Write(textData)
			}, myWindow)
		saveDialog.SetFileName(Tabs.Selected().Text)
		saveDialog.Show()
	})

	myWindow.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("File", fileItem1, fileItem2, fileItem3)))

	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(1000, 600))
	myWindow.ShowAndRun()
}
