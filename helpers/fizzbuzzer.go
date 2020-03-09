package helpers

func Fizzbuzzer(initialNumber int, moduloNumber int, stringToReplace string) string {
	if initialNumber%moduloNumber == 0 {
		return stringToReplace
	}

	return ""
}
