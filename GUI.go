package main

import (
	"log"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var msg = "Server Start at: " + time.Now().Format("Mon Jan _2 15:04:05 2006")

func UI() {

	var mw *walk.MainWindow
	if err := (MainWindow{
		AssignTo: &mw,
		Title:    "Server",
		Layout:   VBox{MarginsZero: true},
		Size:     Size{500, 100},
		Icon:     walk.IconInformation(),
		Children: []Widget{
			Label{
				Text: msg,
			},
		},
		MenuItems: []MenuItem{
			Action{
				Text: "E&xit",
				OnTriggered: func() {
					WriteAll()
					mw.Close()
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	go func() {
		StartWorkers()
		StartGin()
		ConfigRuntime()
	}()

	mw.Run()
}

func Notify() {
	// We need either a walk.MainWindow or a walk.Dialog for their message loop.
	// We will not make it visible in this example, though.
	mw, err := walk.NewMainWindow()
	if err != nil {
		log.Fatal(err)
	}

	// We load our icon from a file.
	icon, err := walk.NewIconFromFile("icon/main.ico")
	if err != nil {
		log.Fatal(err)
	}

	// Create the notify icon and make sure we clean it up on exit.
	ni, err := walk.NewNotifyIcon()
	if err != nil {
		log.Fatal(err)
	}
	defer ni.Dispose()

	// Set the icon and a tool tip text.
	if err := ni.SetIcon(icon); err != nil {
		log.Fatal(err)
	}
	if err := ni.SetToolTip("Click for info or use the context menu to exit."); err != nil {
		log.Fatal(err)
	}

	// When the left mouse button is pressed, bring up our balloon.
	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}

		if err := ni.ShowCustom(
			"Server Running",
			msg); err != nil {

			log.Fatal(err)
		}
	})

	// We put an exit action into the context menu.
	exitAction := walk.NewAction()
	if err := exitAction.SetText("E&xit"); err != nil {
		log.Fatal(err)
	}
	exitAction.Triggered().Attach(func() {
		walk.App().Exit(0)
	})
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		log.Fatal(err)
	}

	// The notify icon is hidden initially, so we have to make it visible.
	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	// Now that the icon is visible, we can bring up an info balloon.
	if err := ni.ShowInfo("Server has just start", msg); err != nil {
		log.Fatal(err)
	}

	// Run the message loop.
	mw.Run()
}
