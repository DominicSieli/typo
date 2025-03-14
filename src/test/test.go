package test

import "fmt"
import "typo/src/input"
import "typo/src/actions"
import "typo/src/terminal"

func Test(text string) {
	index := 0
	correct := 0
	endIndex := 0
	incorrect := 0
	startIndex := 0
	endLines := []int{}
	scoreMap := []bool{}
	startLines := []int{}

	startLines = append(startLines, 0)
	_, lineLimit, err := terminal.TerminalSize()

	if err != nil {
		panic(err)
	}

	lineLimit = lineLimit - 4

	for i := range text {
		if text[i] == 10 || i + 1 >= len(text) {
			endLines = append(endLines, i)
		}

		if text[i] == 10 && i + 1 < len(text) {
			startLines = append(startLines, i + 1)
		}
	}

	if len(startLines) < lineLimit {
		lineLimit = len(startLines)
	}

	endIndex = lineLimit - 1

	for true {
		if (text[index] == 10 || text[index] == 32 || text[index] == 9) && index + 1 < len(text) {
			if text[index] == 10 && startIndex + lineLimit < len(startLines) {
				endIndex++
				startIndex++
			}

			index++
			scoreMap = append(scoreMap, true)
			continue
		}

		render(index, startLines[startIndex], endLines[endIndex], text, scoreMap, correct, incorrect)
		key := input.Key()

		if index >= len(text) - 1 {
			startIndex, endIndex = scroll(key, startIndex, endIndex, len(startLines))
		}

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
			endIndex = lineLimit - 1
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

func scroll(key byte, startIndex int, endIndex int, size int) (int, int) {
	if actions.Up(key) && startIndex > 0 {
		endIndex--
		startIndex--
	}

	if actions.Down(key) && endIndex + 1 < size {
		endIndex++
		startIndex++
	}

	return startIndex, endIndex
}

func render(index int, startIndex int, endIndex int, text string, scoreMap []bool, correct int, incorrect int) {
	terminal.Clear()

	fmt.Printf("%sCorrect: %d%s\n", terminal.GREEN, correct, terminal.RESET)
	fmt.Printf("%sIncorrect: %d%s\n", terminal.RED, incorrect, terminal.RESET)
	fmt.Println()

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
}
