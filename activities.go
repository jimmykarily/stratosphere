package main

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func activityDetails(activity *Activity) *widget.Group {
	generic := widget.NewGroupWithScroller("Values")
	generic.Prepend(widget.NewLabel("Activity Id: " + strconv.Itoa(activity.Id)))
	// TODO: generate image only when user clicks on activity. we don't want to make loading slow
	// Consider caching images
	activity.GenerateMap()
	mapImg := canvas.NewImageFromImage(activity.Image)
	mapImg.SetMinSize(fyne.NewSize(MapWidth, MapHeight))
	generic.Prepend(widget.NewHBox(layout.NewSpacer(), mapImg, layout.NewSpacer()))

	return generic
}

func makeActivitiesList(application App) fyne.Widget {
	activityList := []*widget.TabItem{}

	activities, err := application.db.Activities()
	if err != nil {
		dialog.ShowError(err, application.window)
	}

	for _, activity := range activities {
		activityList = append(activityList,
			widget.NewTabItemWithIcon(activity.Content.StartTimeStr(), nil, activityDetails(activity)),
		)
	}

	activityTabs := widget.NewTabContainer(activityList...)
	activityTabs.SetTabLocation(widget.TabLocationLeading)

	scroll := widget.NewScrollContainer(activityTabs)
	scroll.Resize(fyne.NewSize(200, 200))

	return scroll
}

// WidgetScreen shows a panel containing widget demos
func ActivitiesScreen(application App) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil),
		makeActivitiesList(application),
	)
}
