package render

import "fmt"
import "typo/src/terminal"

func RenderFileList(files []string, index int) {
	terminal.Clear()

	end := 0
	limit := 10
	size := len(files)
	start := (index / limit) * limit

	if start + limit <= size {
		end = start + limit - 1
	}

	if start + limit >= size {
		end = size - 1
	}

	if start > 9 {
		terminal.ColorPrintLine("grey", "<...")
	}

	for i := start; i <= end; i++ {
		if i != index {
			terminal.ColorPrintLine("grey", files[i])
		}

		if i == index {
			terminal.ColorPrintLine("cyan", files[i])
		}
	}

	if end < size - 1 {
		terminal.ColorPrintLine("grey", "...>")
	}

	fmt.Println()
	terminal.ColorPrintLine("grey", "Move up: Up Arrow")
	terminal.ColorPrintLine("grey", "Move down: Down Arrow")
	terminal.ColorPrintLine("grey", "Exit Application: Left Arrow")
	terminal.ColorPrintLine("grey", "Select file: Right Arrow")
}

func RenderText(text string, startIndex int, endIndex int, index int, scoreMap []bool, correct int, incorrect int) {
	terminal.Clear()

	for i := startIndex; i <= endIndex; i++ {
		if i == index {
			terminal.ColorPrintCharacter("cyan", rune(text[i]))
		}

		if i > index {
			terminal.ColorPrintCharacter("grey", rune(text[i]))
		}

		if i < index {
			if scoreMap[i] == true {
				terminal.ColorPrintCharacter("green", rune(text[i]))
			}

			if scoreMap[i] == false {
				terminal.ColorPrintCharacter("red", rune(text[i]))
			}
		}
	}

	fmt.Println()
	fmt.Printf("%sCorrect: %d%s\n", terminal.GREEN, correct, terminal.RESET)
	fmt.Printf("%sIncorrect: %d%s\n", terminal.RED, incorrect, terminal.RESET)
	fmt.Println()
	terminal.ColorPrintLine("grey", "Back to file select: Left Arrow")
}
