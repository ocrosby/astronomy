# Astronomy

[![Lint](https://github.com/ocrosby/astronomy/actions/workflows/lint.yml/badge.svg)](https://github.com/ocrosby/astronomy/actions/workflows/lint.yml)
[![Unittest](https://github.com/ocrosby/astronomy/actions/workflows/test.yml/badge.svg)](https://github.com/ocrosby/astronomy/actions/workflows/test.yml)

An example repo to illustrate some astronomy calculations.

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Development Tasks](#development-tasks)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project provides a set of tools for performing various astronomy calculations.

## Prerequisites

- Go 1.23.0 or later
- [go-task](https://taskfile.dev/) - Install with: `go install github.com/go-task/task/v3/cmd/task@latest`

## Installation

To install all development dependencies, run:

```sh
task deps
```

This will install:
- go-task (build automation)
- golangci-lint (linting)
- Ginkgo v2 (testing framework)
- Gomega (matcher library)

Then install project dependencies:

```sh
task install
```

## Development Tasks

This project uses [go-task](https://taskfile.dev/) for build automation. Available tasks:

```sh
task --list
```

Common tasks:
- `task deps` - Install development dependencies
- `task install` - Install project dependencies  
- `task lint` - Run golangci-lint
- `task test` - Run tests with Ginkgo
- `task clean` - Clean up build artifacts

## Testing

This project uses [Ginkgo](https://onsi.github.io/ginkgo/) v2 and [Gomega](https://onsi.github.io/gomega/) for testing.

Run all tests with coverage:
```sh
task test
```

Run tests for specific packages:
```sh
ginkgo ./pkg/angles
ginkgo ./pkg/solar
```

Run tests with custom options:
```sh
ginkgo -r --randomize-all --fail-on-pending
```

## Contributing

If you would like to contribute, please fork the repository and use a feature branch. Pull requests are welcome.

## References

I wrote these functions after having scanned through the [embedded document from NOAA](./docs/solareqns.pdf).

