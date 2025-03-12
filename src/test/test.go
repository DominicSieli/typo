package test

import "fmt"
import "typo/src/input"
import "typo/src/actions"
import "typo/src/terminal"

var key byte
var index int
var text string
var correct int
var endIndex int
var startIndex int
var incorrect int
var files []string
var scoreMap []bool
var endLines []int
var startLines []int
var lineLimit int = 10

func Test(text string) {
	for true {
		reset()
		populateLines()

		for true {
			scroll()

			if index < len(text) - 1 && (text[index] == 10 || text[index] == 32 || text[index] == 9) {
				scoreMap = append(scoreMap, true)
				key = 0
				index++
				continue
			}

			render()
			key = input.Input()
			update()

			if actions.Escape(key) {
				break
			}

			if actions.Enter(key) {
				reset()
				populateLines()
				continue
			}
		}
	}
}

func reset() {
	key = 0
	index = 0
	correct = 0
	endIndex = 0
	incorrect = 0
	startIndex = 0
	scoreMap = []bool{}
	endLines = []int{}
	startLines = []int{}
}

func scroll() {
	if text[index] == 10 {
		if endIndex < len(endLines) - 1 {
			endIndex++
			startIndex++
		}
	}

	if key == 1 {
		if startIndex > 0 {
			endIndex--
			startIndex--
		}
	}

	if key == 2 {
		if endIndex < len(endLines) - 1 {
			endIndex++
			startIndex++
		}
	}
}

func populateLines() {
	startLines = append(startLines, 0)

	for i := 0; i < len(text); i++ {
		if text[i] == 10 {
			endLines = append(endLines, i)

			if i + 1 < len(text) {
				startLines = append(startLines, i + 1)
			}
		}
	}

	if len(endLines) < lineLimit {
		endIndex = len(endLines) - 1
	}

	if len(endLines) >= lineLimit {
		endIndex = lineLimit - 1
	}
}

func update() {
	if key > 4 {
		if index < len(text) - 1 {
			if key == text[index] {
				scoreMap = append(scoreMap, true)
				index++
				correct++
				return
			}

			if key != text[index] {
				scoreMap = append(scoreMap, false)
				index++
				incorrect++
				return
			}
		}
	}
}

func render() {
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
