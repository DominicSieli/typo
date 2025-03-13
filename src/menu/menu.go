package menu

import "os"
import "fmt"
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
		fmt.Println("No files found")
		fmt.Print(terminal.UNHIDE_CURSOR)
		os.Exit(0)
	}

	for true {
		render(index, files)
		key := input.Key()
		index = actions.Scroll(key, index, len(files))

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
				fmt.Print(terminal.UNHIDE_CURSOR)
				os.Exit(0)
			}

			return text
		}
	}

	return text
}

func render(index int, files []string) {
	end := 0
	size := len(files)
	limit, err := terminal.TerminalHeight()

	if err != nil {
		panic(err)
	}

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
