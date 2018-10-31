package main

import (
	"errors"
	"strconv"
)

var FizzNullError = "fizz value can not be 0"
var BuzzNullError = "buzz value can not be 0"

func GetFizzBuzzString(fizz int, buzz int, limit int, fizzString string, buzzString string) (*[]string, error){
	var result []string

	if fizz == 0 {
		return &result, errors.New(FizzNullError)
	}

	if buzz == 0 {
		return &result, errors.New(BuzzNullError)
	}

	fizzBuzzString := fizzString + buzzString

	for i := 1; i <= limit; i++  {
		if isMultiple(i, fizz) {
			if isMultiple(i, buzz) {
				result = append(result, fizzBuzzString)
			}else {
				result = append(result, fizzString)
			}
		} else if isMultiple(i, buzz) {
			result = append(result, buzzString)
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return &result, nil
}

func isMultiple(val int, base int) (bool) {
	return !(base == 0 || val % base != 0)
}