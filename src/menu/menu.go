package menu

import "os"
import "fmt"
import "typo/src/input"
import "typo/src/fileio"
import "typo/src/actions"
import "typo/src/terminal"

var key byte
var index int
var text string
var files []string
var limit int = 20

func Menu() string {
	for true {
		files = fileio.ReadFiles()

		if len(files) == 0 {
			terminal.Clear()
			fmt.Println("No files found")
			os.Exit(0)
		}

		for true {
			render()
			key := input.Input()
			scroll()

			if actions.Escape(key) {
				terminal.Clear()
				fmt.Print(terminal.UNHIDE_CURSOR)
				os.Exit(0)
			}

			if actions.Enter(key) {
				file := files[index]
				text = fileio.ReadFile(file)

				if len(text) == 0 {
					terminal.Clear()
					fmt.Println("This file is empty")
					os.Exit(0)
				} else {
					return text
				}
			}
		}
	}

	return text
}

func scroll() {
	if actions.Up(key) {
		if index > 0 {
			index--
		}
	}

	if actions.Down(key) {
		if index + limit < len(files) - 1 {
			index++
		}
	}
}

func render() {
	var end int
	size := len(files)
	start := (index / limit) * limit

	terminal.Clear()

	if start + limit < size {
		end = start + limit - 1
	}

	if start + limit >= size {
		end = size - 1
	}

	for i := start; i <= end; i++ {
		if i != index {
			terminal.ColorPrintLine("grey", files[i])
		}

		if i == index {
			terminal.ColorPrintLine("cyan", files[i])
		}
	}

	fmt.Print(terminal.HIDE_CURSOR)
}
