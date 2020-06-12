package main

import (
	"time"

	"github.com/getlantern/systray"
	"github.com/petershen0307/citray/icon"
)

func onReady() {
	systray.SetTemplateIcon(icon.DataGithub, icon.DataGithub)
	systray.SetTitle(time.Now().String())
	systray.SetTooltip("Look at me, I'm a tooltip!")
	testRed := systray.AddMenuItem("test red", "test icon")
	testNormal := systray.AddMenuItem("test normal", "test icon")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")
	go func() {
		for {
			select {
			case <-testNormal.ClickedCh:
				systray.SetIcon(icon.DataGithub)
			case <-testRed.ClickedCh:
				systray.SetIcon(icon.DataGithubRed)
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
	go func() {
		tick := false
		for {
			if tick {
				systray.SetIcon(icon.DataGithub)
			} else {
				systray.SetIcon(icon.DataGithubRed)
			}
			tick = !tick
			time.Sleep(1 * time.Second)
		}
	}()
}
