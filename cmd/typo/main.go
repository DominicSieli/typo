package main

import "typo/internal/menu"
import "typo/internal/test"
import "typo/internal/terminal"

func main() {
	terminal.Clear()

	for true {
		text := menu.Menu()
		test.Test(text)
	}

	terminal.Clear()
}
