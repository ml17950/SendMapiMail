@echo off

set GOARCH=386
go build -o SendMapiMail32.exe .

set GOARCH=amd64
go build -o SendMapiMail64.exe .
