package terminal

import "os"
import "fmt"
import "os/exec"

const (
	RESET = "\033[0m"
	RED = "\033[38;2;255;0;0m"
	GREEN = "\033[38;2;0;255;0m"
	CYAN = "\033[38;2;0;255;255m"
	GREY = "\033[38;2;128;128;128m"
)

func ColorPrintLine(color string, text string) {
	switch color {
	case "red":
		color = RED
	case "green":
		color = GREEN
	case "cyan":
		color = CYAN
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
	case "cyan":
		color = CYAN
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
