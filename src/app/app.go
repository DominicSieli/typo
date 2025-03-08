package app

import "os"
import "fmt"
import "typo/src/input"
import "typo/src/fileio"
import "typo/src/render"
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

func Loop() {
	for true {
		Reset()
		files = fileio.PopulateFileList()

		if len(files) == 0 {
			terminal.Clear()
			fmt.Println("No files found")
			os.Exit(0)
		}

		for true {
			render.RenderFileList(files, index)
			key = input.RawInput()
			NavigateFileList()

			if actions.Escape(key) {
				terminal.Clear()
				os.Exit(0)
			}

			if actions.Enter(key) {
				file := files[index]
				text = fileio.ReadFile(file)
				break
			}
		}

		Reset()
		PopulateLines()

		for true {
			ScrollText()

			if index < len(text) - 1 && (text[index] == 10 || text[index] == 32 || text[index] == 9) {
				scoreMap = append(scoreMap, true)
				key = 0
				index++
				continue
			}

			render.RenderText(text, startLines[startIndex], endLines[endIndex], index, scoreMap, correct, incorrect)
			key = input.RawInput()
			UpdateText()

			if actions.Escape(key) {
				break
			}

			if actions.Enter(key) {
				Reset()
				PopulateLines()
				continue
			}
		}
	}
}

func Reset() {
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

func ScrollText() {
	if text[index] == 10 {
		if endIndex < len(endLines) - 1 {
			endIndex++
			startIndex++
		}
	}

	if key == 65 {
		if startIndex > 0 {
			endIndex--
			startIndex--
		}
	}

	if key == 66 {
		if endIndex < len(endLines) - 1 {
			endIndex++
			startIndex++
		}
	}
}

func PopulateLines() {
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

func NavigateFileList() {
	if key == 65 {
		if index > 0 {
			index--
		}
	}

	if key == 66 {
		if index < (len(files) - 1) {
			index++
		}
	}
}

func UpdateText() {
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
