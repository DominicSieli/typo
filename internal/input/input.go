package input

import "os"
import "golang.org/x/term"

func Key() byte {
	buffer := [3]byte{}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	_, err = os.Stdin.Read(buffer[:])

	if err != nil {
		panic(err)
	}

	if buffer[0] == 0x1B && buffer[1] == 0x5B {
		switch buffer[2] {
		case 0x41:
			return 1
		case 0x42:
			return 2
		case 0x43:
			return 3
		case 0x44:
			return 4
		default:
			return 0
		}
	}

	return buffer[0]
}
