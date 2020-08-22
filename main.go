package main

import (
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

var w fyne.Window
var a fyne.App

const key = "Marker"

func main() {
	if os.Args[1] == "new" {
		New(os.Args[2])
	} else if os.Args[1] == "find" {
		Find()
	}
}

func menus(win fyne.Window) {
	lightButton := fyne.NewMenuItem("Light Theme", func() {
		a.Settings().SetTheme(theme.LightTheme())
	})
	darkButton := fyne.NewMenuItem("Dark Theme", func() {
		a.Settings().SetTheme(theme.DarkTheme())
	})
	themeMenu := fyne.NewMenu("Theme", lightButton, darkButton)

	mainMenu := fyne.NewMainMenu(themeMenu)
	win.SetMainMenu(mainMenu)
}
