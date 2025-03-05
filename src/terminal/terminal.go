package terminal

import "os"
import "fmt"
import "os/exec"

const (
	RED = "\033[38;2;255;0;0m"
	GREEN = "\033[38;2;0;255;0m"
	YELLOW = "\033[38;2;255;255;0m"
	BLUE = "\033[38;2;0;0;255m"
	MAGENTA = "\033[38;2;255;0;255m"
	CYAN = "\033[38;2;0;255;255m"
	WHITE = "\033[38;2;255;255;255m"
	GREY = "\033[38;2;200;200;200m"
	RESET = "\033[0m"
)

func ColorPrintLine(color string, text string) {
	switch color {
	case "red":
		color = RED
	case "green":
		color = GREEN
	case "yellow":
		color = YELLOW
	case "blue":
		color = BLUE
	case "magenta":
		color = MAGENTA
	case "cyan":
		color = CYAN
	case "white":
		color = WHITE
	case "grey":
		color = GREY
	default:
		color = ""
	}

	fmt.Println(color + text + RESET)
}

func ColorPrintCharacter(color string, character rune) {
	switch color {
	case "red":
		color = RED
	case "green":
		color = GREEN
	case "yellow":
		color = YELLOW
	case "blue":
		color = BLUE
	case "magenta":
		color = MAGENTA
	case "cyan":
		color = CYAN
	case "white":
		color = WHITE
	case "grey":
		color = GREY
	default:
		color = ""
	}

	fmt.Printf("%s%c%s", color, character, RESET)
}

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
