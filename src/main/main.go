package main

import "typo/src/app"
import "typo/src/terminal"

func main() {
	terminal.Clear()

	app.Loop()

	terminal.Clear()
}
