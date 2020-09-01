package main

import (
	"encoding/json"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var searchBox fyne.CanvasObject

// Find show a dialog with markers
func Find() {
	a = app.NewWithID("com.nvcode.marker")
	w = a.NewWindow("View Markers")
	menus(w)

	Show("")

	w.Resize(fyne.NewSize(500, 500))

	w.ShowAndRun()
}

// Show shows the existing ones in a list of buttons
func Show(search string) {
	existing := []map[string]string{}
	json.Unmarshal([]byte(a.Preferences().StringWithFallback(key, "[]")), &existing)

	if searchBox == nil {
		searchBox = newSearchEntry()
		searchBox.(*searchEntry).SetPlaceholder("Search...")
	}

	btns := []fyne.CanvasObject{
		searchBox,
	}

	for i, val := range existing {
		if search == "" {
			capture := val
			capture["Pos"] = strconv.Itoa(i)
			btns = append(btns, widget.NewButton(capture["Title"], func() { View(capture) }))
		} else {
			if strings.Contains(val["Title"], search) {
				capture := val
				capture["Pos"] = strconv.Itoa(i)
				btns = append(btns, widget.NewButton(capture["Title"], func() { View(capture) }))
			}
		}
	}

	vbox := widget.NewVBox(btns...)
	scroll := widget.NewScrollContainer(vbox)

	w.SetContent(scroll)
}

type searchEntry struct {
	widget.Entry
}

func newSearchEntry() *searchEntry {
	entry := &searchEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (s *searchEntry) KeyDown(key *fyne.KeyEvent) {
	s.Entry.KeyDown(key)
	go Show(s.Entry.Text)
}

func (s *searchEntry) SetPlaceholder(placeholder string) {
	s.Entry.PlaceHolder = placeholder
}

func (s *searchEntry) SetText(text string) {
	s.Entry.SetText(text)
}

func (s *searchEntry) SetCursor(column int) {
	s.Entry.CursorColumn = column
}

func (s *searchEntry) Focus() {
	s.Entry.FocusGained()
}
