# Online Code Execution Tool

A HTTP-based code execution service written in Go, using language-specific Podman containers for secure and isolated runtime environments. Supports Python, C, C++, and Go execution.

## Requirements

- `Go`
- `Podman`

## Commands

```sh
make build # setup and build the executable
make run # runs the executable
make clean # deletes the executable and logs
make test # runs the test curl scripts
```

## API Usage

Sample request at `/`

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
