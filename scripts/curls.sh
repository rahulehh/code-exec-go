#!/bin/bash

# test python
curl -X POST http://localhost:4202/ \
     -H "Content-Type: application/json" \
     -d '{"language":"python","code":"print(\"Hello from Python\")"}'

# test go
curl -X POST http://localhost:4202/ \
     -H "Content-Type: application/json" \
     -d '{"language":"go","code":"package main\nimport \"fmt\"\nfunc main() {\n    fmt.Println(\"Hello from Go\")\n}"}'

# test c
curl -X POST http://localhost:4202/ \
     -H "Content-Type: application/json" \
     -d '{"language":"c","code":"#include <stdio.h>\nint main() {\n    printf(\"Hello from C\\n\");\n    return 0;\n}"}'

# test cpp
curl -X POST http://localhost:4202/ \
     -H "Content-Type: application/json" \
     -d '{"language":"cpp","code":"#include <iostream>\nint main() {\n    std::cout << \"Hello from C++\" << std::endl;\n    return 0;\n}"}'
