package helpers

import (
	"net/url"
	"testing"
)

func TestGetIntParams(t *testing.T) {
	queryValues := url.Values{
		"keyString": []string{"value"},
		"keyInt":    []string{"10"},
	}
	expectedInt := 10

	_, err := GetIntParams(queryValues, "wrongKey")
	if err == nil {
		t.Errorf("Error missing int key test: expected an error")
	}

	if err.Error() != ErrorMissingParams {
		t.Errorf("Error missing int key test : expected %s got %s", ErrorMissingParams, err.Error())
	}

	_, err = GetIntParams(queryValues, "keyString")
	if err == nil {
		t.Errorf("Error wrong type int params test: expected an error")
	}

	if err.Error() != ErrorWrongTypeParams {
		t.Errorf("Error wrong type int params test : expected %s got %s", ErrorWrongTypeParams, err.Error())
	}

	keyInt, err := GetIntParams(queryValues, "keyInt")
	if err != nil {
		t.Fatal(err)
	}

	if keyInt != expectedInt {
		t.Errorf("Error get int params test : expected %d got %d", expectedInt, keyInt)
	}
}

func TestGetStringParams(t *testing.T) {
	queryValues := url.Values{
		"keyString": []string{"value"},
	}
	expectedString := "value"

	_, err := GetStringParams(queryValues, "wrongKey")
	if err == nil {
		t.Errorf("Error missing string key test: expected an error")
	}

	if err.Error() != ErrorMissingParams {
		t.Errorf("Error missing string key test : expected %s got %s", ErrorMissingParams, err.Error())
	}

	keyString, err := GetStringParams(queryValues, "keyString")
	if err != nil {
		t.Fatal(err)
	}

	if keyString != expectedString {
		t.Errorf("Error get string params test : expected %s got %s", expectedString, keyString)
	}
}
