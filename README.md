# Online Code Execution Tool

A HTTP-based code execution service written in Go, using language-specific Docker containers for secure and isolated runtime environments. Supports Python, C, C++, and Go execution.

## Requirements

- `Go`
- `Docker`

## Setup

Before running the service, make sure to run the following script to build all the required images:

```bash
./setup.sh
```

## Running and Building

To build the service:

```bash
go build -o code-exec ./cmd/server/
```

To run the application (make sure Docker is running...)

```bash
go run ./cmd/server/main.go
# or
./code-exec # If the application is built
```

## API Usage

```json
{
  "language": "python",
  "code": "print('Hello, World!')"
}
```

Expected response:

```json
{
  "stdout": "Hello, World!\n",
  "stderr": "",
  "error": ""
}
```

Supported values for `language`:

- `python`
- `go`
- `c`
- `cpp`

Additionally, all supported languages can be tested by running:

```bash
./scripts/curls.sh
```
