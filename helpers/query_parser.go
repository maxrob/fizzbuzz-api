package helpers

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

func GetIntParams(queryValues url.Values, paramsName string) (int, error) {
	params, ok := queryValues[paramsName]

	if !ok || len(params[0]) < 1 {
		return 0, errors.New(fmt.Sprintf("missing-params", paramsName))
	}

	param, err := strconv.Atoi(params[0])
	if err != nil {
		return 0, errors.New(fmt.Sprintf("wrong-type-params"))
	}

	return param, nil
}

func GetStringParams(queryValues url.Values, paramsName string) (string, error) {
	params, ok := queryValues[paramsName]

	if !ok || len(params[0]) < 1 {
		return "", errors.New(fmt.Sprintf("missing-params", paramsName))
	}

	return params[0], nil
}
