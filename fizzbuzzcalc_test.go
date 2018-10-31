package main

import (
	"strings"
	"testing"
)

func TestIsMultipleTrue(t *testing.T) {
	result := isMultiple(28, 7)
	if result != true {
		t.Errorf("IsMultiple was incorrect, got: %t, want: %t.", result, true)
	}
}

func TestIsMultipleFalse(t *testing.T) {
	result := isMultiple(29, 7)
	if result != false {
		t.Errorf("IsMultiple was incorrect, got: %t, want: %t.", result, false)
	}
}

func TestIsMultipleError(t *testing.T) {
	result := isMultiple(28, 0)
	if result != false {
		t.Errorf("IsMultiple was incorrect, got: %t, want: %t.", result, false)
	}
}

func TestIsMultiple0(t *testing.T) {
	result := isMultiple(0, 4)
	if result != true {
		t.Errorf("IsMultiple was incorrect, got: %t, want: %t.", result, true)
	}
}

func TestGetFizzBuzzString(t *testing.T)  {
	expected := []string{"1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz","16","17","fizz","19","buzz"}
	result, _ := GetFizzBuzzString(3, 5, 20, "fizz", "buzz")

	expectedString := strings.Join(expected, ",")
	resultString := strings.Join(*result, ",")

	if resultString != expectedString {
		t.Errorf("FizzBuzzString was incorrect, got: %s, want: %s", resultString, expectedString)
	}
}

func TestGetFizzBuzzStringI(t *testing.T)  {
	expected := []string{"1","2","buzz","4","fizz","buzz","7","8","buzz","fizz","11","buzz","13","14","fizzbuzz","16","17","buzz","19","fizz"}
	result, _ := GetFizzBuzzString(5, 3, 20, "fizz", "buzz")

	expectedString := strings.Join(expected, ",")
	resultString := strings.Join(*result, ",")

	if resultString != expectedString {
		t.Errorf("FizzBuzzString was incorrect, got: %s, want: %s", resultString, expectedString)
	}
}

func TestGetFizzBuzzLimit0(t *testing.T)  {
	expected := []string{}
	result, _ := GetFizzBuzzString(3, 5, 0, "fizz", "buzz")

	expectedString := strings.Join(expected, ",")
	resultString := strings.Join(*result, ",")

	if resultString != expectedString {
		t.Errorf("FizzBuzzString was incorrect, got: %s, want: %s", resultString, expectedString)
	}
}

func TestGetFizzBuzzLowLimit(t *testing.T)  {
	expected := []string{"1", "2", "3", "4", "5"}
	result, _ := GetFizzBuzzString(8, 9, 5, "fizz", "buzz")

	expectedString := strings.Join(expected, ",")
	resultString := strings.Join(*result, ",")

	if resultString != expectedString {
		t.Errorf("FizzBuzzString was incorrect, got: %s, want: %s", resultString, expectedString)
	}
}

func TestGetFizzBuzzFizz1(t *testing.T)  {
	expected := []string{"fizz", "fizzbuzz", "fizz", "fizzbuzz", "fizz"}
	result, _ := GetFizzBuzzString(1, 2, 5, "fizz", "buzz")

	expectedString := strings.Join(expected, ",")
	resultString := strings.Join(*result, ",")

	if resultString != expectedString {
		t.Errorf("FizzBuzzString was incorrect, got: %s, want: %s", resultString, expectedString)
	}
}

func TestGetFizzBuzzFizz0(t *testing.T)  {
	_, err := GetFizzBuzzString(0, 2, 5, "fizz", "buzz")

	if err == nil {
		t.Errorf("FizzBuzzString was incorrect, got: %s, want: %s", "nil", FizzNullError )
	} else if err.Error() != FizzNullError {
		t.Errorf("FizzBuzzString was incorrect, got: %s, want: %s", err.Error(), FizzNullError )
	}
}

func TestGetFizzBuzzBuzz0(t *testing.T)  {
	_, err := GetFizzBuzzString(1, 0, 5, "fizz", "buzz")

	if err == nil {
		t.Errorf("FizzBuzzString was incorrect, got: %s, want: %s", "nil", BuzzNullError )
	} else if err.Error() != BuzzNullError {
		t.Errorf("FizzBuzzString was incorrect, got: %s, want: %s", err.Error(), BuzzNullError )
	}
}

