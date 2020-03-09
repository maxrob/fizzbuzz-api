package helpers

import (
	"errors"
	"net/url"
	"strconv"
)

func GetIntParams(queryValues url.Values, paramsName string) (int, error) {
	params, ok := queryValues[paramsName]

	if !ok || len(params[0]) < 1 {
		return 0, errors.New(ErrorMissingParams)
	}

	param, err := strconv.Atoi(params[0])
	if err != nil {
		return 0, errors.New(ErrorWrongTypeParams)
	}

	return param, nil
}

func GetStringParams(queryValues url.Values, paramsName string) (string, error) {
	params, ok := queryValues[paramsName]

	if !ok || len(params[0]) < 1 {
		return "", errors.New(ErrorMissingParams)
	}

	return params[0], nil
}
