package actions

func Enter(key byte) bool {
	return key == 3
}

func Escape(key byte) bool {
	return key == 4
}
