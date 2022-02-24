package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/youla-dev/godenv/ast"
	"github.com/youla-dev/godenv/parser"
	"github.com/youla-dev/godenv/scanner"
)

func TestParser_Parse(t *testing.T) {
	t.Parallel()

	t.Run("parse assigment successful", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			expected ast.Statement
		}{
			{
				name:  "unquoted value",
				input: "name=value",
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "name",
							Value: "value",
						},
					},
				},
			},
			{
				name:  "double quoted value",
				input: `name="value"`,
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "name",
							Value: "value",
						},
					},
				},
			},
			{
				name:  "single quoted value",
				input: `name='value'`,
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "name",
							Value: "value",
						},
					},
				},
			},
			{
				name:  "name with assign and empty value",
				input: "name=",
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "name",
							Value: "",
						},
					},
				},
			},
			{
				name:  "name without value",
				input: "name",
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "name",
							Value: "",
						},
					},
				},
			},
			{
				name:  "variable with blank lines",
				input: "\n\n\n\nname=\n\n\n",
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "name",
							Value: "",
						},
					},
				},
			},
			{
				name:  "multiple variables",
				input: "DEBUG_HTTP_ADDR=:9090\nDEBUG_HTTP_IDLE_TIMEOUT=0s\nJAEGER_AGENT_ENDPOINT=jaeger-otlp-agent:6831",
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "DEBUG_HTTP_ADDR",
							Value: ":9090",
						},
						&ast.AssignStatement{
							Name:  "DEBUG_HTTP_IDLE_TIMEOUT",
							Value: "0s",
						},
						&ast.AssignStatement{
							Name:  "JAEGER_AGENT_ENDPOINT",
							Value: "jaeger-otlp-agent:6831",
						},
					},
				},
			},
			{
				name:  "variable with comments",
				input: "# comment 1\nDEBUG_HTTP_ADDR=:9090\n# comment 2",
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.CommentStatement{
							Value: "# comment 1",
						},
						&ast.AssignStatement{
							Name:  "DEBUG_HTTP_ADDR",
							Value: ":9090",
						},
						&ast.CommentStatement{
							Value: "# comment 2",
						},
					},
				},
			},

			{
				name:  "newlines in quoted strings",
				input: `FOO="bar\nbaz"`,
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "FOO",
							Value: "bar\\nbaz",
						},
					},
				},
			},
			{
				name:  "single quotes inside double quotes",
				input: `FOO="'d'"`,
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "FOO",
							Value: "'d'",
						},
					},
				},
			},
			{
				name:  `variable with several "=" in the value`,
				input: `FOO=foobar=`,
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "FOO",
							Value: "foobar=",
						},
					},
				},
			},
			{
				// FIXME (i-sevostyanov): should be fixed in the future
				name:  `inline comments is a part of value`,
				input: `FOO=bar # this is foo`,
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "FOO",
							Value: "bar # this is foo",
						},
					},
				},
			},
			{
				name:  `allows # in double quoted value`,
				input: `FOO="bar#baz"`,
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "FOO",
							Value: "bar#baz",
						},
					},
				},
			},
			{
				name:  `allows # in single quoted value`,
				input: `FOO='bar#baz'`,
				expected: &ast.FileStatement{
					Statements: []ast.Statement{
						&ast.AssignStatement{
							Name:  "FOO",
							Value: "bar#baz",
						},
					},
				},
			},
		}

		for _, tt := range tests {
			tt := tt

			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				s := scanner.New(tt.input)
				p := parser.New(s)

				stmts, err := p.Parse()
				require.NoError(t, err)
				assert.Equal(t, tt.expected, stmts)
			})
		}
	})

	t.Run("returns error on invalid input", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{
				name:  "escaped double quotes",
				input: `FOO="escaped\"bar"`,
			},
			{
				name:  "value with space after equal sign",
				input: `FOO= bar`,
			},
			{
				name:  "value with space before equal sign",
				input: `FOO =bar`,
			},
			{
				name:  "leading tab",
				input: "\tFOO=bar",
			},
			{
				name:  "leading whitespace",
				input: "  FOO=bar",
			},
		}

		for _, tt := range tests {
			tt := tt

			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				s := scanner.New(tt.input)
				p := parser.New(s)

				stmts, err := p.Parse()
				require.Error(t, err)
				assert.Nil(t, stmts)
			})
		}
	})
}
