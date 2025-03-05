package render

import "fmt"
import "typo/src/terminal"

func RenderFileList(files []string, index int) {
	terminal.Clear()

	for i, v := range files {
		if i != index {
			terminal.ColorPrintLine("grey", v)
		}

		if i == index {
			terminal.ColorPrintLine("cyan", v)
		}
	}
}

func RenderText(text string, index int, scoreMap []bool, correct int, incorrect int) {
	terminal.Clear()

	for i, v := range text {
		if i == index {
			terminal.ColorPrintCharacter("cyan", v)
		}

		if i > index {
			terminal.ColorPrintCharacter("grey", v)
		}

		if i < index {
			if scoreMap[i] == true {
				terminal.ColorPrintCharacter("green", v)
			}

			if scoreMap[i] == false {
				terminal.ColorPrintCharacter("red", v)
			}
		}
	}

	fmt.Println()
	fmt.Printf("%sCorrect: %d%s\n", terminal.GREEN, correct, terminal.RESET)
	fmt.Printf("%sIncorrect: %d%s\n", terminal.RED, incorrect, terminal.RESET)
}
