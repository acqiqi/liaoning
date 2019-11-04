#!/bin/bash
GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags="-s -w" ./main.go