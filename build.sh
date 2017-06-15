#!/bin/bash

echo "Compiling for Windows"
GOOS=windows GOARCH=amd64 go build -o caa-twit.exe main.go
echo "Compiling for Linux"
GOOS=linux GOARCH=amd64 go build -o caa-twit.linux main.go
echo "Compiling for MacOS"
GOOS=darwin GOARCH=amd64 go build -o caa-twit.mac main.go
