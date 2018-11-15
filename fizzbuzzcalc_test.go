package main

import (
	"strings"
	"testing"
)

func TestGetFizzBuzzString(t *testing.T)  {
	expected := []string{"1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz","16","17","fizz","19","buzz","fizz","22","23","fizz","buzz","26","fizz","28","29","fizzbuzz","31","32","fizz","34","buzz","fizz","37","38","fizz","buzz","41","fizz","43","44","fizzbuzz","46","47","fizz","49","buzz","fizz","52","53","fizz","buzz","56","fizz","58","59","fizzbuzz","61","62","fizz","64","buzz","fizz","67","68","fizz","buzz","71","fizz","73","74","fizzbuzz","76","77","fizz","79","buzz","fizz","82","83","fizz","buzz","86","fizz","88","89","fizzbuzz","91","92","fizz","94","buzz","fizz","97","98","fizz","buzz"}
	result, _ := GetFizzBuzzString(3, 5, 100, "fizz", "buzz")

	expectedString := strings.Join(expected, ",")
	resultString := strings.Join(*result, ",")

	if resultString != expectedString {
		t.Errorf("FizzBuzzString was incorrect, got: %s, want: %s", resultString, expectedString)
	}
}

func TestGetFizzBuzzStringI(t *testing.T)  {
	expected := []string{"1","2","buzz","4","fizz","buzz","7","8","buzz","fizz","11","buzz","13","14","fizzbuzz","16","17","buzz","19","fizz","buzz","22","23","buzz","fizz","26","buzz","28","29","fizzbuzz","31","32","buzz","34","fizz","buzz","37","38","buzz","fizz","41","buzz","43","44","fizzbuzz","46","47","buzz","49","fizz","buzz","52","53","buzz","fizz","56","buzz","58","59","fizzbuzz","61","62","buzz","64","fizz","buzz","67","68","buzz","fizz","71","buzz","73","74","fizzbuzz","76","77","buzz","79","fizz","buzz","82","83","buzz","fizz","86","buzz","88","89","fizzbuzz","91","92","buzz","94","fizz","buzz","97","98","buzz","fizz"}
	result, _ := GetFizzBuzzString(5, 3, 100, "fizz", "buzz")

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

