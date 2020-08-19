package main

import (
	"encoding/json"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var title *widget.Entry
var content *widget.Entry
var p string

// New brings up the new menu
func New(path string) {
	p = path

	a = app.NewWithID("com.nvcode.marker")
	w = a.NewWindow("New Marker")
	menus(w)

	title = widget.NewEntry()
	content = widget.NewMultiLineEntry()

	form := widget.NewForm()
	form.Append("Title", title)
	form.Append("Description", content)

	submit := widget.NewButtonWithIcon("Create", theme.ConfirmIcon(), Submit)
	cancel := widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {
		w.Close()
		a.Quit()
	})

	hbox := widget.NewHBox(layout.NewSpacer(), cancel, submit, layout.NewSpacer())
	vbox := widget.NewVBox(form, hbox)
	scroll := widget.NewScrollContainer(vbox)

	w.Resize(fyne.NewSize(300, 200))

	w.SetContent(scroll)
	w.ShowAndRun()
}

// Submit adds the marker to the preferences
func Submit() {
	existing := []map[string]string{}
	json.Unmarshal([]byte(a.Preferences().StringWithFallback("Markers", "[]")), &existing)
	existing = append(existing, map[string]string{"Path": p, "Title": title.Text, "Content": content.Text})
	prefs, _ := json.Marshal(existing)
	a.Preferences().SetString("Markers", string(prefs))

	w.Close()
	a.Quit()
}
