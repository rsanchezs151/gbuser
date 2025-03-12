@echo off
echo Batch Script to take input.
set /p input= Escriba el comentario:  
git add .
git commit -m "%input%"
git push

set GOOS=linux
set GOARCH=amd64
go build -tags lambda.norpc -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip main