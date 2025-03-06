package actions

func Escape(key byte) bool {
	return key == 68
}

func Enter(key byte) bool {
	return key == 67
}
