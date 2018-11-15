package main

import (
	"errors"
	"strconv"
)

var FizzNullError = "fizz value can not be 0"
var BuzzNullError = "buzz value can not be 0"
var LimitExceededError = "limit value can not exceed 10000000"

func GetFizzBuzzString(fizz int, buzz int, limit int, fizzString string, buzzString string) (*[]string, error){
	var result []string

	if limit > 10000000 {
		return &result, errors.New(LimitExceededError)
	}

	if fizz == 0 {
		return &result, errors.New(FizzNullError)
	}

	if buzz == 0 {
		return &result, errors.New(BuzzNullError)
	}

	fizzBuzzString := fizzString + buzzString

	multFizz := fizz
	multBuzz := buzz
	fizzBuzz := fizz * buzz
	multFizzBuzz := fizzBuzz
	for i := 1; i <= limit; i++  {
		if i == multFizzBuzz {
			result = append(result, fizzBuzzString)
			multFizzBuzz = multFizzBuzz + fizzBuzz
			multFizz = multFizz + fizz
			multBuzz = multBuzz + buzz
		} else if i == multBuzz {
			result = append(result, buzzString)
			multBuzz = multBuzz + buzz
		} else if i == multFizz {
			result = append(result, fizzString)
			multFizz = multFizz + fizz
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return &result, nil
}
