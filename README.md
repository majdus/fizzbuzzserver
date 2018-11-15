# fizzbuzzServer

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by “fizz”, all multiples of 5 by “buzz”, and all multiples of 15 by “fizzbuzz”. The output would look like this:
“1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...”

This version of fizzbuss expose a REST API endpoint
that accepts five parameters :
    - three integers (int1, int2 and limit)
    - and two strings (string1 and string2),
and returns a JSON

It must return a list of strings with numbers from 1 to limit, where:
all multiples of int1 are replaced by string1,
all multiples of int2 are replaced by string2,
all multiples of int1 and int2 are replaced by string1string2

## Run on localhost
### Get and install dependencies
    go get -d -v
    go install -v
### Build
    go build
### Run
    $GOBIN/fizzbuzzserver

## Run on docker
### Build image
    docker build -t fizzbuzz .
### Run container
    docker run --rm --publish 8000:8000 -t -i fizzbuzz

## Try it
[local deployment with int1 = 3, int2 = 50, limit = 20, string1 = "fizz", string2 = "buzz"](http://localhost:8000/3/5/20/fizz/buzz)

## Try it on now.sh running server
[Docker deployed server](https://fizzbuzzserver-eaqrludlsn.now.sh/int1/3/int2/5/limit/20/string1/fizz/string2/buzz)
