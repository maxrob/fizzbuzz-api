package helpers

func Fizzbuzzer(initialNumber int, multipleNumber int, stringToReplace string) string {
	if initialNumber%multipleNumber == 0 {
		return stringToReplace
	}

	return ""
}
