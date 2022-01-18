# Specification

This file contains the complete `godenv` specification.

## CHARACTER SET

All `.env` files MUST be in the **UTF-8** encoding.

## Syntax

### COMMENTS

If the current line is started with a `#`, the line is interpreted as a comment. The `<newline>` character is excluded from the comment. Empty lines are ignored.

```dotenv
# Comment example. It starts from the '#' symbol and ends at the new line.

# ^ the empty line above is ignored. 
```

### NAME=VALUE

Any line `<name>=<value>` is treated as an environment variable assignment.
- The expression at the left of the `=` symbol is the variable name.
- The expression at the right of the `=` symbol is the variable value.

Whitespace characters (` `, `\t`, etc.) on the left of the `=` symbol are restricted and lead to the scan error.

`<name>` is a token that MUST be composed of Unicode letters, Unicode digits, and the following characters: `_` (underscore), `,` (comma), `.` (full stop), `-` (hyphen).
The name MUST contain at least 1 character.
- Unicode letters mean the characters defined in the Letter categories Lu, Ll, Lt, Lm, or Lo of [The Unicode Standard 8.0](https://www.unicode.org/versions/Unicode8.0.0/).
- Unicode digits mean the characters defined in the Number category Nd of [The Unicode Standard 8.0](https://www.unicode.org/versions/Unicode8.0.0/).

`<value>` is a UTF-8 string. It can be quoted or unquoted. Its interpretation slightly differs depending on the quotes used.
The value MUST be a single-line string.
- A value in single quotes. The value is a string between `'` characters. Both quotation marks are excluded from the value.

  All characters inside a single-quoted value are treated as-is and are not escaped.
- A value in double-quotes. The value is a string between `"` characters. Both quotation marks are excluded from the value.

  The special character sequences are escaped. E.g., `\n` is converted into `<newline>`.
- An unquoted value string. The value is a string between `=` character and a `<newline>`. Both `=` and `<newline>` are excluded from the value.

  The special character sequences are escaped. E.g., `\n` is converted into `<newline>`.

```dotenv
# Valid variables
valid-name.with_special,symbols=value
КИРИЛЛИЦА_IS_SUPPRTED_AS_WELL=value
value_without_quotes=A value without quotes will be interpreted as a value.
value_with_single_quotes='A value between single quotation marks.'
value_with_double_quotes="A value between double quotation marks."

# The following names are implied illegal:
#  - this/name/contains/slashes
#  - LEGAL_NAME_BUT_WITH_SPACE =value

# The following values are implied illegal:
# 
# VALUE_WITHOUT_CLOSING_QUOTE="Illegal end of the line.
# 
# MULTI_LINE_VALUE="Multi line values
# are not supported yet"
#
# ILLEGAL_ESCAPE_SEQUENCE="\ <- this slash MUST be escaped as '\\'."
```

#### SPECIAL CASES

- If a value is empty, it's interpreted as an empty string ''.
- If a line lacks `=` character, the value is interpreted as an empty string ''.
- If a file contains 2 or more equal `<name>`-s, then the last value is used.

```dotenv
# Value is an empty string ''
VARIABLE_WITH_EMPTY_VALUE=
# Value is an empty string '' 
VARIABLE_WITH_NO_EQUAL_CHAR

# Duplicated variables
VAR_NAME=value1
# The result is "value2" 
VAR_NAME=value2
```