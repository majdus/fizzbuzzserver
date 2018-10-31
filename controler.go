package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

var Fizz = "fizz"
var Buzz = "buzz"
var Limit = "limit"
var FizzString = "fizzString"
var BuzzString = "buzzString"

type FizzBuzzRequest struct {
	Fizz int
	Buzz int
	Limit int
	FizzString string
	BuzzString string
}

type FizzBuzzResponse struct {
	Result *[]string	`json:"result"`
}

type FizzBuzzResponseError struct {
	Message string	`json:"message"`
}

func Usage(w http.ResponseWriter, r *http.Request) {
	response := FizzBuzzResponseError{
		Message: "Example : " +
			"[http://localhost:8000/int1/3/int2/5/limit/20/string1/fizz/string2/buzz] or " +
			"[http://localhost:8000/3/5/20/fizz/buzz]",
	}
	SendJSON(w, response)
}

func GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request, err := getRequestData(params)
	if err != nil {
		response := FizzBuzzResponseError{
			Message: err.Error(),
		}
		SendJSON(w, response)
		return
	}

	fizzBuzzString, err := GetFizzBuzzString(request.Fizz, request.Buzz, request.Limit, request.FizzString, request.BuzzString)
	if err != nil {
		response := FizzBuzzResponseError{
			Message: err.Error(),
		}
		SendJSON(w, response)
		return
	}

	response := FizzBuzzResponse{
		Result: fizzBuzzString,
	}
	SendJSON(w, response)
}

func getRequestData(params map[string]string) (*FizzBuzzRequest, error) {
	var request FizzBuzzRequest
	var err error

	request.Fizz, err = strconv.Atoi(params[Fizz])
	if err != nil {
		return nil, errors.New("fizz param error, fizz = " + params[Fizz] + " - error : " + err.Error())
	}
	request.Buzz, err = strconv.Atoi(params[Buzz])
	if err != nil {
		return nil, errors.New("buzz param error, buzz = " + params[Buzz] + " - error : " + err.Error())
	}
	request.Limit, err = strconv.Atoi(params[Limit])
	if err != nil {
		return nil, errors.New("limi param error, limit = " + params[Limit] + " - error : " + err.Error())
	}
	request.FizzString = params[FizzString]
	request.BuzzString = params[BuzzString]

	return  &request, nil
}

// SendJSON writes a struct to the writer
func SendJSON(w http.ResponseWriter, i interface{}) {
	js, err := json.Marshal(i)
	if err != nil {
		http.Error(w, "JSON Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func WaitForCtrlC() {
	var ctrlCWaiter sync.WaitGroup
	ctrlCWaiter.Add(1)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		ctrlCWaiter.Done()
	}()

	ctrlCWaiter.Wait()
}
