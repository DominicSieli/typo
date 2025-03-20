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
	startLines := []int{}
	scoreMap := map[byte]int{}

	startLines = append(startLines, 0)
	_, lineLimit, err := terminal.TerminalSize()

	if err != nil {
		panic(err)
	}

	lineLimit = lineLimit - 6

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
			endIndex = lineLimit - 1
			scoreMap = map[byte]int{}
			continue
		}

		if key > 4 && index + 1 < len(text) {
			if key == text[index] {
				key = 0
				index++
				correct++
				continue
			}

			if key != text[index] {
				key = 0
				incorrect++

				if _, k := scoreMap[text[index]]; !k {
					scoreMap[text[index]] = 1
					continue
				}

				if v, k := scoreMap[text[index]]; k {
					scoreMap[text[index]] = v + 1
					continue
				}
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

func render(index int, startIndex int, endIndex int, text string, scoreMap map[byte]int, correct int, incorrect int) {
	terminal.Clear()

	fmt.Printf("%sCorrect: %d%s\n", terminal.GREEN, correct, terminal.RESET)
	fmt.Printf("%sIncorrect: %d%s\n", terminal.RED, incorrect, terminal.RESET)

	for key, value := range scoreMap {
		fmt.Printf(terminal.RED + "[%c:%d]" + terminal.RESET, rune(key), value)
	}

	fmt.Println()
	fmt.Println()

	for i := startIndex; i <= endIndex; i++ {
		if i < index {
			terminal.ColorPrintCharacter("green", rune(text[i]))
		}

		if i == index {
			terminal.ColorPrintCharacter("cyan", rune(text[i]))
		}

		if i > index {
			terminal.ColorPrintCharacter("grey", rune(text[i]))
		}
	}
}
