package actions

func Up(key byte) bool {
	return key == 1
}

func Down(key byte) bool {
	return key == 2
}

func Enter(key byte) bool {
	return key == 3 || key == 13
}

func Escape(key byte) bool {
	return key == 4 || key == 27
}

func Scroll(key byte, index int, size int) int {
	if Up(key) && index > 0 {
		index--
	}

	if Down(key) && index + 1 < size {
		index++
	}

	return index
}
