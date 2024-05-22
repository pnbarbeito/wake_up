package main

//go:generate go-winres make

import (
	"fmt"
	"wake_up/icon"
	"wake_up/move"
	"github.com/getlantern/systray"
)

func main() {
	onExit := func() {
		fmt.Println("Exit")
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Wake Up!")
	systray.SetTooltip("Wake Up!")
	mQuitOrig := systray.AddMenuItem("Salir", "Cerrar Wake Up!")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	go func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
		systray.SetTitle("Wake Up!")
		systray.SetTooltip("Wake Up!")
		move.Move_win()
	}()
}
