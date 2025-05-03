#!/bin/sh 

chmod 777 /app
chmod 777 /tmp
cat - > /app/code.go
chmod +x /app/code.go
go run /app/code.go
