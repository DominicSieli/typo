package actions

func Escape(key byte) bool {
	return key == 127
}

func Enter(key byte) bool {
	return key == 13
}
