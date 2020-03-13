package main

import (
	"fmt"
	"time"

	"github.com/alexrett/idler"
	"github.com/getlantern/systray"
)

type DB struct {
	idle int
}

var db = DB{idle: 0}

func loopStart() {
	i := idler.NewIdle()
	for {
		db.idle = i.GetIdleTime()
		time.Sleep(1 * time.Second)
		systray.SetTitle(fmt.Sprintf("IDLE %d", db.idle))
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
	systray.SetTitle(fmt.Sprintf("IDLE %d", db.idle))
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
