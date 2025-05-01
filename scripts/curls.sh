#!/bin/bash

# Make the curl POST request with valid JSON
curl --unix-socket /tmp/code-exec.sock \
     -X POST http://localhost/ \
     -H "Content-Type: application/json" \
     -d '{"language":"python","code":"print(Hello World)"}'
