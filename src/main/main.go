package main

import "typo/src/menu"
import "typo/src/test"
import "typo/src/terminal"

func main() {
	terminal.Clear()

	for true {
		text := menu.Menu()
		test.Test(text)
	}

	terminal.Clear()
}
