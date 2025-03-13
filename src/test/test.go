package test

import "fmt"
import "typo/src/input"
import "typo/src/actions"
import "typo/src/terminal"

func Test(text string) {
	index := 0
	correct := 0
	incorrect := 0
	startIndex := 0
	scoreMap := []bool{}

	for true {
		if (text[index] == 10 || text[index] == 32 || text[index] == 9) && index + 1 < len(text) {
			index++
			scoreMap = append(scoreMap, true)
			continue
		}

		render(index, startIndex, text, scoreMap, correct, incorrect)
		key := input.Key()
		// startIndex = actions.Scroll(key, startIndex, len(text))

		if actions.Escape(key) {
			break
		}

		if actions.Enter(key) {
			key = 0
			index = 0
			correct = 0
			incorrect = 0
			startIndex = 0
			scoreMap = []bool{}
			continue
		}

		if key > 4 && index + 1 < len(text) {
			if key == text[index] {
				key = 0
				index++
				correct++
				scoreMap = append(scoreMap, true)
				continue
			}

			if key != text[index] {
				key = 0
				index++
				incorrect++
				scoreMap = append(scoreMap, false)
				continue
			}
		}
	}
}

func render(index int, startIndex int, text string, scoreMap []bool, correct int, incorrect int) {
	terminal.Clear()

	fmt.Printf("%sCorrect: %d%s\n", terminal.GREEN, correct, terminal.RESET)
	fmt.Printf("%sIncorrect: %d%s\n", terminal.RED, incorrect, terminal.RESET)
	fmt.Println()

	for i := startIndex; i < len(text); i++ {
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
}
