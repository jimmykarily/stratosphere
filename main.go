package main

import (
	"fmt"
	"log"
	"net/url"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var db *Db

type App struct {
	app    fyne.App
	db     *Db
	window fyne.Window
}

func welcomeScreen(a fyne.App) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(theme.FyneLogo())
	logo.SetMinSize(fyne.NewSize(128, 128))

	link, err := url.Parse("https://fyne.io/")
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return widget.NewVBox(
		widget.NewLabelWithStyle("Welcome to the Fyne toolkit demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		widget.NewHyperlinkWithStyle("fyne.io", link, fyne.TextAlignCenter, fyne.TextStyle{}),
		layout.NewSpacer(),

		widget.NewGroup("Theme",
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				widget.NewButton("Dark", func() {
					a.Settings().SetTheme(theme.DarkTheme())
				}),
				widget.NewButton("Light", func() {
					a.Settings().SetTheme(theme.LightTheme())
				}),
			),
		),
	)
}

func main() {
	var err error
	var application App

	application.db, err = NewDb("fixtures/test.sqlite3")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer application.db.connection.Close()

	application.app = app.New()
	application.window = application.app.NewWindow("Stratosphere - free software for your watch")
	application.window.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() { fmt.Println("Menu New") }),
		// a quit item will be appended to our first menu
	), fyne.NewMenu("Edit",
		fyne.NewMenuItem("Cut", func() { fmt.Println("Menu Cut") }),
		fyne.NewMenuItem("Copy", func() { fmt.Println("Menu Copy") }),
		fyne.NewMenuItem("Paste", func() { fmt.Println("Menu Paste") }),
	)))

	tabs := widget.NewTabContainer(
		//widget.NewTabItemWithIcon("Welcome", theme.HomeIcon(), welcomeScreen(a)),
		widget.NewTabItemWithIcon("Activities", theme.ContentCopyIcon(), ActivitiesScreen(application)),
		widget.NewTabItemWithIcon("Widgets", theme.ContentCopyIcon(), WidgetScreen()),
		widget.NewTabItemWithIcon("Graphics", theme.DocumentCreateIcon(), GraphicsScreen()),
		widget.NewTabItemWithIcon("Windows", theme.ViewFullScreenIcon(), DialogScreen(application.window)),
		widget.NewTabItemWithIcon("Advanced", theme.SettingsIcon(), AdvancedScreen(application.window)),
	)
	tabs.SetTabLocation(widget.TabLocationLeading)
	application.window.SetContent(tabs)

	application.window.ShowAndRun()
}
