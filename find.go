package main

import (
	"encoding/json"
	"os/exec"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
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
	json.Unmarshal([]byte(a.Preferences().StringWithFallback("Markers", "[]")), &existing)

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

// View is for editing a value
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
			json.Unmarshal([]byte(a.Preferences().StringWithFallback("Markers", "[]")), &existing)
			pos, _ := strconv.Atoi(val["Pos"])

			copy(existing[pos:], existing[pos+1:])
			existing[len(existing)-1] = make(map[string]string, 0)
			existing = existing[:len(existing)-1]

			prefs, _ := json.Marshal(existing)
			a.Preferences().SetString("Markers", string(prefs))

			Show("")
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
