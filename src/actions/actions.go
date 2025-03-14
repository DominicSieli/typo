package actions

func Up(key byte) bool {
	return key == 1
}

func Down(key byte) bool {
	return key == 2
}

func Enter(key byte) bool {
	return key == 13
}

func Escape(key byte) bool {
	return key == 27
}
