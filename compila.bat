@echo off
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -tags lambda.norpc -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip main bootstrap