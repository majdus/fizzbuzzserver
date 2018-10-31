#!/bin/sh -e
docker build -t fizzbuzz .
docker run --rm --publish 8000:8000 -t -i fizzbuzz
