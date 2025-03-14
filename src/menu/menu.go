package menu

import "os"
import "fmt"
import "time"
import "typo/src/input"
import "typo/src/fileio"
import "typo/src/actions"
import "typo/src/terminal"

func Menu() string {
	text := ""
	index := 0
	files := fileio.ReadFiles()

	if len(files) == 0 {
		terminal.Clear()
		terminal.ColorPrintLine("red", "[NO FILES FOUND]")
		fmt.Print(terminal.UNHIDE_CURSOR)
		os.Exit(0)
	}

	for true {
		render(index, files)
		key := input.Key()
		index = scroll(key, index, len(files))

		if actions.Escape(key) {
			terminal.Clear()
			fmt.Print(terminal.UNHIDE_CURSOR)
			os.Exit(0)
		}

		if actions.Enter(key) {
			file := files[index]
			text = fileio.ReadFile(file)

			if len(text) == 0 {
				index = 0
				terminal.Clear()
				terminal.ColorPrintLine("red", "[THIS FILE IS EMPTY]")
				time.Sleep(time.Duration(3) * time.Second)
				continue
			}

			return text
		}
	}

	return text
}

func scroll(key byte, index int, size int) int {
	if actions.Up(key) && index > 0 {
		index--
	}

	if actions.Down(key) && index + 1 < size {
		index++
	}

	return index
}

func render(index int, files []string) {
	end := 0
	size := len(files)

	terminal.Clear()

	_, height, err := terminal.TerminalSize()

	if err != nil {
		panic(err)
	}

	if index + height <= size {
		end = index + height - 2
	}

	if index + height > size {
		end = size
	}

	for i := index; i < end; i++ {
		if i != index {
			terminal.ColorPrintLine("grey", files[i])
		}

		if i == index {
			terminal.ColorPrintLine("cyan", files[i])
		}
	}

	fmt.Print(terminal.HIDE_CURSOR)
}
