#!/bin/bash

export DYNAMO_ENDPOINT=http://localhost:8000/

go mod tidy
go run main/sample.go