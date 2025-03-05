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
var incorrect int
var files []string
var scoreMap []bool

func Loop() {
	for true {
		Reset()
		files = fileio.PopulateFileList()

		if len(files) == 0 {
			terminal.Clear()
			fmt.Println("no files found")
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

		for true {
			render.RenderText(text, index, scoreMap, correct, incorrect)
			key = input.RawInput()
			key = input.TranslateEnterKey(key)

			if index == len(text) - 1 {
				Reset()
				continue
			}

			if actions.Escape(key) {
				break
			}

			UpdateText()
		}
	}
}

func Reset() {
	key = 0
	index = 0
	correct = 0
	incorrect = 0
	scoreMap = []bool{}
}

func NavigateFileList() {
	if key == 66 {
		if index < (len(files) - 1) {
			index++
		}
	}

	if key == 65 {
		if index > 0 {
			index--
		}
	}
}

func UpdateText() {
	if index < len(text) {
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
