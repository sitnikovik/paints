# Paints 🎨

[![Go Reference](https://pkg.go.dev/badge/github.com/sitnikovik/paints.svg)](https://pkg.go.dev/github.com/sitnikovik/paints)
[![Go Report Card](https://goreportcard.com/badge/github.com/sitnikovik/paints)](https://goreportcard.com/report/github.com/sitnikovik/paints)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sitnikovik/paints)
[![Release](https://img.shields.io/github/v/release/sitnikovik/paints?style=flat)](https://github.com/sitnikovik/paints/releases)

Toolkit to prettify the output of your golang apps

## Features

- Simple functions to colorize text and backgrounds
- Text styles: bold, dim, italic, underline
- Zero runtime dependencies — pure Go implementation
- Small, well-tested functions suitable for CLI output

## Installation

```go
go get github.com/sitnikovik/paints
```

## Getting started

```go
package main

import (
    "fmt"

    bg "github.com/sitnikovik/paints/color/background"
    color "github.com/sitnikovik/paints/color/text"
    style "github.com/sitnikovik/paints/style/text"
)

func main() {
    // Print text with red background
    s := bg.Red("Lorem ipsum")
    fmt.Println(s)

    // Print green bold text
    s = color.Green(style.Bold("Lorem ipsum"))
    fmt.Println(s)
}
```

## Documentation

Full API reference is available on [pkg.go.dev](https://pkg.go.dev/github.com/sitnikovik/paints)


## Project structure

- `color/` — public packages `text` and `background` for coloring strings
- `style/` — text style functions (bold, italic, etc.)
- `internal/ansi/` — ANSI escape codes and helpers (internal)
- `docs/` — documentation and release notes
- Tests are colocated with packages (files ending with `_test.go`)

## Requirements

- Go 1.23.4 (as declared in `go.mod`) — newer Go versions should work
- No external runtime dependencies

## Contributing

Contributions are welcome. Suggested workflow:

1. Pick an open issue and work on it. Open a Pull Request that references the issue when you ready
2. If there is no suitable issue for your idea, either open a short issue describing the proposal or contact me (watch my [bio](https://github.com/sitnikovik)) and we will create an issue for you.
3. Fork the repository and create a feature branch: `git checkout -b feature/name`
4. Run and add tests for new behavior: `go test ./...`
5. Commit changes with a descriptive message and open a Pull Request
6. Follow the repository's code style and linters (we use `golangci-lint`)
7. If the PR passes, we merge it and create version via Github releases

Please open an issue to discuss large or breaking changes before implementing.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

## Author / Contact

Maintained by [Ilya Sitnikov](https://github.com/sitnikovik)
