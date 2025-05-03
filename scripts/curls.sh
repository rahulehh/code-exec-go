#!/bin/bash

# test python
curl -X POST http://localhost:4202/ \
     -H "Content-Type: application/json" \
     -d '{"language":"python","code":"print(\"Hello from Python\")"}'

# test go
curl -X POST http://localhost:4202/ \
     -H "Content-Type: application/json" \
     -d '{"language":"go","code":"package main\nimport \"fmt\"\nfunc main() {\n    fmt.Println(\"Hello from Go\")\n}"}'
