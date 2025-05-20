#!/bin/sh

chmod 777 /app
chmod 777 /tmp

cat - > /app/code.c

gcc /app/code.c -o /app/code.out 2> /app/compile_error.txt

if [ -f /app/code.out ]; then
    /app/code.out
else
    cat /app/compile_error.txt
fi
