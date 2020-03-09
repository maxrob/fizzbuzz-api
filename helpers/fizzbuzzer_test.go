package helpers

import "testing"

func TestFizzbuzzer(t *testing.T) {
	moduloNumber := 3
	stringToReplace := "fizz"

	result1 := Fizzbuzzer(1, moduloNumber, stringToReplace)
	if result1 != "" {
		t.Errorf("Wrong result1, got: %s, want: %s.", result1, "")
	}

	result3 := Fizzbuzzer(3, moduloNumber, stringToReplace)
	if result3 != stringToReplace {
		t.Errorf("Wrong result3, got: %s, want: %s.", result3, stringToReplace)
	}

	result15 := Fizzbuzzer(15, moduloNumber, stringToReplace)
	if result15 != stringToReplace {
		t.Errorf("Wrong result15, got: %s, want: %s.", result15, stringToReplace)
	}

	result20 := Fizzbuzzer(20, moduloNumber, stringToReplace)
	if result20 != "" {
		t.Errorf("Wrong result20, got: %s, want: %s.", result20, "")
	}
}
