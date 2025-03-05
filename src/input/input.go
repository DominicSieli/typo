package input

import "os"
import "golang.org/x/term"

func RawInput() byte {
	var buf [1]byte

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	_, err = os.Stdin.Read(buf[:])

	if err != nil {
		panic(err)
	}

	return buf[0]
}

func TranslateEnterKey(key byte) byte {
	if key == 13 {
		return 10
	}

	return key
}
