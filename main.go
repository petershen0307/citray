package main

import "github.com/getlantern/systray"

func main() {
	onExit := func() {
	}

	systray.Run(onReady, onExit)
}
