package main

import (
	"fmt"
	"time"

	"github.com/alexrett/active-window"
	"github.com/alexrett/idler"
	"github.com/getlantern/systray"
)

type DB struct {
	idle         int
	activeWindow string
}

var db = DB{idle: 0, activeWindow: ""}

func loopStart() {
	i := idler.NewIdle()
	a := activeWindow.ActiveWindow{}
	for {
		db.idle = i.GetIdleTime()
		db.activeWindow, _ = a.GetActiveWindowTitle()
		systray.SetTitle(fmt.Sprintf("%s %d", db.activeWindow, db.idle))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	onExit := func() {
		fmt.Println("Finished onExit")
	}
	// Should be called at the very beginning of main().

	systray.RunWithAppWindow("Lantern", 1024, 768, onReady, onExit)
}

func onReady() {
	systray.SetTitle(fmt.Sprintf("%s %d", db.activeWindow, db.idle))
	systray.SetTooltip("Lantern")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	go loopStart()
}
