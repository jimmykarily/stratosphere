package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// Callbacks must be closures or your callback will use the last value
// of the activity variable in the for loop!
// https://tour.golang.org/moretypes/25
func loadActivityCallback(activity *Activity) func() {
	return func() {
		fmt.Println(activity.Content)
		fmt.Println(activity.Id)
	}
}

func makeActivitiesList(application App) fyne.Widget {
	logo := canvas.NewImageFromResource(theme.FyneLogo())
	logo.SetMinSize(fyne.NewSize(320, 320))
	list := widget.NewVBox()
	activities, err := application.db.Activities()
	if err != nil {
		dialog.ShowError(err, application.window)
	}

	for _, activity := range activities {
		list.Append(widget.NewButton(activity.Content.StartTimeStr(), loadActivityCallback(activity)))
	}

	scroll := widget.NewScrollContainer(list)
	scroll.Resize(fyne.NewSize(200, 200))
	container := widget.NewHBox(scroll)

	return container
}

// WidgetScreen shows a panel containing widget demos
func ActivitiesScreen(application App) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil),
		makeActivitiesList(application),
	)
}
