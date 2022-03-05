# godenv — a proper package to read `.env` files

[The Twelve-Factor App Manifest] promotes using environment variables as the only configuration
approach. We follow and love this principle.

The .env files are just a genuine extension of env vars. The .env files simplify local debug,
exposing variables to a Docker containers, and so on.

Godenv is a tiny package that reads those .env files.

## Motivation

We took inspiration from the [godotenv] repository. The goal we pursued was to write a parser
without using regular expressions but with a lexer/parser approach.

If you are curious about learning more about the approach, see the following links:

- https://en.wikipedia.org/wiki/Abstract_syntax_tree
- https://en.wikipedia.org/wiki/Lexical_analysis
- https://en.wikipedia.org/wiki/Parsing#Parser

## Installation

```shell
go get github.com/youla-dev/godenv
```

## Usage

Add a configuration to your `.env` file:

```dotenv
HTTP_ADDRESS=":8080"
LOG_LEVEL="info"
```

Then in the Go app read the file and parse its content:

```go
package main

import (
	"fmt"
	"os"

	"github.com/youla-dev/godenv"
)

func main() {
	f, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	vars, err := godenv.Parse(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(vars)
}
```

## Pronunciation

`godenv` stands for go-dot-env. It is pronounced as `goh denv`, not as `gahd env`.

## Specification

The complete specification can be found here: [SPECIFICATION.md](SPECIFICATION.md).

[//]: # (Links)

[godotenv]: https://github.com/joho/godotenv

[The Twelve-Factor App Manifest]: https://12factor.net/
