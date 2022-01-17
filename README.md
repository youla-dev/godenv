# godenv — a proper package to read `.env` files

## Pronunciation

`godenv` stands for go-dot-env. It is pronounced as `goh denv`, not as `gahd env`.  

## Specification

```dotenv
# CHARACTER SET: All `.env` files MUST be in the UTF-8 encoding.

# COMMENT: If the current line is started with a '#', the line is interpreted as a comment. The <newline> character is excluded from the comment.
# Empty lines are ignored.

# NAME=VALUE
# Any line in the form <name>=<value> is treated as environment variable assignment.
# The expression at the left of the '=' symbol is the variable name.
# The expression at the right of the '=' symbol is the variable value.
#
# Whitespace characters (' ', '\t', etc.) on the left of the '=' symbol are restricted and lead to the scan error.
# 
# <name> is a token that must be composed of Unicode letters, Unicode digits, 
# and the following characters: '_' (underscore), ',' (comma), '.' (full stop), '-' (hyphen).
# The name MUST contain at least 1 character.
#  - Unicode letters mean the characters defined in the Letter categories Lu, Ll, Lt, Lm, or Lo of The Unicode Standard 8.0 
#    (https://www.unicode.org/versions/Unicode8.0.0/).
#  - Unicode digits mean the characters defined in the Number category Nd of The Unicode Standard 8.0.
#
# <value> is a UTF-8 string. It can be quoted or unquoted. Its interpretation slightly differs depending on the used quotes.
# The value MUST be a single-line string.
#  - An unquoted value string. The value is a string between '=' character and a <newline>. Both '=' and <newline> are excluded from the value.
#  - A value in single quotes. The value is a string between "'" characters. Both quotation marks are excluded from the value.
#    All characters inside a single-quoted value are treated as-is and are not escaped.
#  - A value in double-quotes. The value is a string between '"' characters. Both quotation marks are excluded from the value.
#    The special character sequences are escaped. E.g., `\n` is converted into <newline>.
_valid.name_with-all-possible,CHARACTERS1=value

value_without_quotes=A value without quotes will be interpreted as a value.
value_with_single_quotes='A value between single quotation marks. Any special character will be treated as-is and will not be escaped.'
value_with_double_quotes="A value between double quotation marks. All special characters (like '\n') will be escaped."

# The following names are implied illegal:
#  - this/name/contains/slashes
#  - LEGAL_NAME_BUT_WITH_SPACE =value

# SPECIAL CASES
# If a value is empty, it's interpreted as an empty string ''.
VARIABLE_WITH_EMPTY_VALUE=
# If a line lacks '=' character, the value is interpreted as an empty string ''. 
VARIABLE_WITH_NO_EQUAL_CHAR
```
