@echo off

set GOARCH=amd64
set GOOS=linux
set CGO_ENABLED=0
set GO111MODULE=on

go build -ldflags="-w -s" -o server .
