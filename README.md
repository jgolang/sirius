# sirius Package

## Overview

The `sirius` package provides a set of functions for working with JSON files in Go. It allows developers to open JSON files and retrieve different types of data (string, int, float, bool, etc.) from JSON keys, facilitating better configuration management and data extraction.

## Installation

To install the `sirius` package, use the following command:

```bash
go get github.com/yourusername/sirius
```

## Usage

Here’s a basic example of how to use the `sirius` package:

### Example JSON File (`config.json`)

```json
{
    "app": {
        "name": "SiriusApp",
        "version": "1.0.0"
    },
    "server": {
        "port": 8080,
        "enabled": true
    },
    "limits": {
        "timeout": 30.5,
        "retries": 5
    }
}
```

### Example Code

```go
package main

import (
    "fmt"
    "log"
    "github.com/yourusername/sirius"
)

func main() {
    config, err := sirius.Open("config.json")
    if err != nil {
        log.Fatalf("Failed to open config file: %v", err)
    }

    appName := config.GetString("app.name")
    port := config.GetInt("server.port")
    timeout := config.GetFloat64("limits.timeout")
    retries := config.GetInt("limits.retries")
    enabled := config.GetBool("server.enabled")

    fmt.Printf("App Name: %s\n", appName)
    fmt.Printf("Port: %d\n", port)
    fmt.Printf("Timeout: %.2f\n", timeout)
    fmt.Printf("Retries: %d\n", retries)
    fmt.Printf("Server Enabled: %t\n", enabled)
}
```

## Functions

### `Open`

```go
func Open(name string) (*SJson, error)
```

Opens a JSON file and parses its contents.

- **name**: The name of the JSON file.
- **Returns**: An `SJson` instance and an error if the file cannot be opened or parsed.

### `Get`

```go
func (g SJson) Get(key string) interface{}
```

Gets a JSON node from a JSON key.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The value associated with the JSON key as an `interface{}`.

### `GetString`

```go
func (g SJson) GetString(key string) string
```

Gets a string JSON node from a JSON key.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The string value associated with the JSON key.

### `GetInt`

```go
func (g SJson) GetInt(key string) int
```

Gets an integer JSON node from a JSON key.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The int value associated with the JSON key.

### `GetInt64`

```go
func (g SJson) GetInt64(key string) int64
```

Gets an integer 64 JSON node from a JSON key.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The int64 value associated with the JSON key.

### `GetFloat64`

```go
func (g SJson) GetFloat64(key string) float64
```

Gets a float 64 JSON node from a JSON key.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The float64 value associated with the JSON key.

### `GetBool`

```go
func (g SJson) GetBool(key string) bool
```

Gets a boolean JSON node from a JSON key.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The boolean value associated with the JSON key.

### `GetArrayMaps`

```go
func (g SJson) GetArrayMaps(key string) []map[string]interface{}
```

Gets an array JSON node from a JSON key. Returns a JSON array as a slice of maps.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The JSON array as a slice of maps.

## Structs

### `SJson`

```go
type SJson struct {
    Buf  []byte
    node jsonNode
}
```

The `SJson` struct holds the buffer and the parsed JSON node.

## Contributing

If you have suggestions for how We could be improved, or want to report a bug, open an issue! We'd love all and any contributions.

For more, check out the [Contributing Guide](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](LICENSE).

## Support

If you find this repository helpful and would like to support its development, consider making a donation. Your contributions will help ensure the continued improvement and maintenance of this repository.

Thank you for your support!

[![ko-fi](https://www.ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/josuegiron)