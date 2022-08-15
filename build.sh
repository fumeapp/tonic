#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o main main.go && zip -r main.zip main config