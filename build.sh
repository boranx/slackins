#!/bin/bash

echo "install dependencies"
go get gopkg.in/jarcoal/httpmock.v1
go get github.com/sbstjn/hanu

echo "format - test - build"
go test -v && go fmt && go build .
