package main

import (
	"encoding/json"
	"os/exec"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// View is for viewing a value
func View(val map[string]string) {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() {
			Show("")
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			cmd := exec.Command("open", "-R", val["Path"])
			cmd.Start()
		}),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			existing := []map[string]string{}
			json.Unmarshal([]byte(a.Preferences().StringWithFallback(key, "[]")), &existing)
			pos, _ := strconv.Atoi(val["Pos"])

			copy(existing[pos:], existing[pos+1:])
			existing[len(existing)-1] = make(map[string]string, 0)
			existing = existing[:len(existing)-1]

			prefs, _ := json.Marshal(existing)
			a.Preferences().SetString(key, string(prefs))

			Show("")
		}),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			Edit(val)
		}),
	)

	title := widget.NewLabel(val["Title"])
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	path := widget.NewLabel(val["Path"])
	path.Alignment = fyne.TextAlignCenter

	content := widget.NewLabel(val["Content"])
	content.Alignment = fyne.TextAlignCenter

	vbox := widget.NewVBox(toolbar, title, path, content)
	scroll := widget.NewScrollContainer(vbox)

	w.SetContent(scroll)
}

// Edit shows the screen for editing a value
func Edit(val map[string]string) {
	title = widget.NewEntry()
	title.PlaceHolder = "Title..."

	content = widget.NewMultiLineEntry()
	content.PlaceHolder = "Description..."

	form := widget.NewForm()
	form.Append("Title", title)
	form.Append("Description", content)

	cancel := widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() { View(val) })
	submit := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		existing := []map[string]string{}
		json.Unmarshal([]byte(a.Preferences().StringWithFallback(key, "[]")), &existing)
		pos, _ := strconv.Atoi(val["Pos"])
		existing[pos] = map[string]string{"Path": p, "Title": title.Text, "Content": content.Text}
		prefs, _ := json.Marshal(existing)
		a.Preferences().SetString(key, string(prefs))
		View(map[string]string{"Path": p, "Title": title.Text, "Content": content.Text, "Pos": val["Pos"]})
	})

	hbox := widget.NewHBox(layout.NewSpacer(), cancel, submit, layout.NewSpacer())
	vbox := widget.NewVBox(form, hbox)
	scroll := widget.NewScrollContainer(vbox)

	w.SetContent(scroll)
}
