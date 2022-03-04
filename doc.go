// Copyright 2022 The godenv authors
// Licensed under the MIT License
//
// https://opensource.org/licenses/MIT
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package godenv is a tiny module that parses .env files.
//
// Motivation
//
// We took inspiration from the godotenv (https://github.com/joho/godotenv)
// repository. The goal we pursued was to write a parser without using regular
// expressions but with a lexer/parser approach.
//
// The current specification of .env files format is listed in SPECIFICATION.md
// (https://github.com/youla-dev/godenv/blob/main/SPECIFICATION.md).
//
// If you are curious about learning more about the approach, see the following links:
//
//   - https://en.wikipedia.org/wiki/Abstract_syntax_tree
//   - https://en.wikipedia.org/wiki/Lexical_analysis
//   - https://en.wikipedia.org/wiki/Parsing#Parser
//
// Usage
//
// Let's assume, you have a .env file with the following content:
//
// 	HTTP_LISTEN=":8080"
// 	LOG_LEVEL="info"
//
// You can easily open the file and parse its content into map[string]string:
//
// 	f, err := os.Open(".env")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
//
// 	vars, err := godenv.Parse(f)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(vars)
//
package godenv
