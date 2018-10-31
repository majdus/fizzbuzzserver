package main

import "testing"

func TestGetRequesData(t *testing.T) {
	var params map[string]string
	params = make(map[string]string)

	params[Fizz] = "4"
	params[Buzz] = "7"
	params[Limit] = "100"
	params[FizzString] = "fizz"
	params[BuzzString] = "buzz"

	result, err := getRequestData(params)

	if err != nil  {
		t.Errorf("getRequesData was incorrect, got: not nil error : %s", err.Error())
	}

	if result.Fizz != 4  {
		t.Errorf("getRequesData was incorrect, got: fizz = %d, want: %d.", result.Fizz, 4)
	}

	if result.Buzz != 7  {
		t.Errorf("getRequesData was incorrect, got: buzz = %d, want: %d.", result.Buzz, 7)
	}

	if result.Limit != 100  {
		t.Errorf("getRequesData was incorrect, got: limit = %d, want: %d.", result.Limit, 100)
	}

	if result.FizzString != "fizz"  {
		t.Errorf("getRequesData was incorrect, got: fizzString = %s, want: %s.", result.FizzString, "fizz")
	}

	if result.BuzzString != "buzz"  {
		t.Errorf("getRequesData was incorrect, got: buzzString = %s, want: %s.", result.BuzzString, "buzz")
	}

}
