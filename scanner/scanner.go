package scanner

import (
	"strconv"
	"unicode"
	"unicode/utf8"

	"github.com/youla-dev/godenv/token"
)

const (
	bom = 0xFEFF // byte order mark, only permitted as the first character
	eof = -1     // eof indicates the end of the file.
)

// Scanner converts a sequence of characters into a sequence of tokens.
type Scanner struct {
	input      string
	ch         rune // current character
	prevOffset int  // position before current character
	offset     int  // character offset
	peekOffset int  // position after current character
}

// New returns new Scanner.
func New(input string) *Scanner {
	s := &Scanner{input: input}

	s.next()
	if s.ch == bom {
		s.next() // ignore BOM at the beginning of the file
	}

	return s
}

// NextToken scans the next token and returns the token position, the token, and its literal string
// if applicable. The source end is indicated by token.EOF.
//
// If the returned token is a literal (token.Identifier, token.Value, token.RawValue) or token.Comment,
// the literal string has the corresponding value.
//
// If the returned token is token.Illegal, the literal string is the offending character.
func (s *Scanner) NextToken() token.Token {
	switch s.ch {
	case eof:
		return token.New(token.EOF, s.offset)
	case '\n':
		return s.scanNewLine()
	case ' ', '\t', '\r', '\v', '\f':
		defer s.next()
		return token.NewWithLiteral(token.Space, string(s.ch), s.offset)
	case '=':
		defer s.next()
		return token.New(token.Assign, s.offset)
	case '#':
		return s.scanComment()
	case '"':
		return s.scanQuotedValue(token.Value, s.ch)
	case '\'':
		return s.scanQuotedValue(token.RawValue, s.ch)
	default:
		switch prev := s.prev(); prev {
		case '\n', bom:
			if isValidIdentifier(s.ch) {
				return s.scanIdentifier()
			}
		case '=':
			return s.scanUnquotedValue()
		}
		return s.scanIllegalRune()
	}
}

// ========================================================================
// Methods that scan a specific token kind.
// ========================================================================

func (s *Scanner) scanNewLine() token.Token {
	for isNewLine(s.ch) {
		s.next()
	}

	return token.NewWithLiteral(token.NewLine, "\n", s.offset)
}

func (s *Scanner) scanIdentifier() token.Token {
	start := s.offset

	for isLetter(s.ch) || isDigit(s.ch) || isSymbol(s.ch) {
		s.next()
	}

	literal := s.input[start:s.offset]

	return token.NewWithLiteral(token.Identifier, literal, s.offset)
}

func (s *Scanner) scanComment() token.Token {
	start := s.offset

	for !(isEOF(s.ch) || isNewLine(s.ch)) {
		s.next()
	}

	lit := s.input[start:s.offset]

	return token.NewWithLiteral(token.Comment, lit, s.offset)
}

func (s *Scanner) scanIllegalRune() token.Token {
	literal := string(s.ch)
	offset := s.offset
	s.next()

	return token.NewWithLiteral(token.Illegal, literal, offset)
}

func (s *Scanner) scanUnquotedValue() token.Token {
	start := s.offset

	for !isEOF(s.ch) && !isNewLine(s.ch) {
		s.next()
	}

	lit := s.input[start:s.offset]

	return token.NewWithLiteral(token.Value, lit, s.offset)
}

func (s *Scanner) scanQuotedValue(tType token.Type, quote rune) token.Token {
	// opening quote already consumed
	s.next()
	start := s.offset

	for {
		if isEOF(s.ch) || isNewLine(s.ch) {
			// TODO (titusjaka): return human-readable error instead
			tType = token.Illegal
			break
		}
		if s.ch == quote {
			break
		}
		s.next()
	}

	offset := s.offset
	lit := s.input[start:offset]

	if s.ch == quote {
		s.next()
	}

	return token.NewWithLiteral(tType, lit, offset)
}

// ========================================================================
// Methods that control pointers to the current, previous, and next chars.
// ========================================================================

// Read the next Unicode char into s.ch.
// s.ch < 0 means end-of-file.
func (s *Scanner) next() {
	s.prevOffset = s.offset

	if s.peekOffset < len(s.input) {
		s.offset = s.peekOffset
		r, width := s.scanRune(s.offset)

		s.peekOffset += width
		s.ch = r
	} else {
		s.offset = len(s.input)
		s.ch = eof
	}

	if s.offset == 0 {
		s.prevOffset = -1
	}
}

func (s *Scanner) prev() rune {
	switch {
	case s.prevOffset < 0:
		return '\n'
	case s.prevOffset < len(s.input):
		r, _ := s.scanRune(s.prevOffset)
		return r
	default:
		return eof
	}
}

// Reads a single Unicode character and returns the rune and its width in bytes.
func (s *Scanner) scanRune(offset int) (r rune, width int) {
	r = rune(s.input[offset])
	width = 1

	switch {
	case r >= utf8.RuneSelf:
		// not ASCII
		r, width = utf8.DecodeRune([]byte(s.input[offset:]))
		if r == utf8.RuneError && width == 1 {
			panic("illegal UTF-8 encoding on position " + strconv.Itoa(offset))
		} else if r == bom && s.offset > 0 {
			panic("illegal byte order mark on position " + strconv.Itoa(offset))
		}
	}
	return r, width
}

// ========================================================================
// Auxiliary methods that check if the rune is one of the specific kind.
// ========================================================================

func isValidIdentifier(r rune) bool {
	return isLetter(r) || isDigit(r) || isSymbol(r)
}

func isLetter(r rune) bool {
	return ('a' <= lower(r) && lower(r) <= 'z') ||
		(r >= utf8.RuneSelf && unicode.IsLetter(r))
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isSymbol(r rune) bool {
	switch r {
	case '_', '.', ',', '-':
		return true
	}
	return false
}

func isNewLine(r rune) bool {
	return r == '\n'
}

func isEOF(r rune) bool {
	return r == eof
}

// ------------------------------------------------------------------------

func lower(r rune) rune { return ('a' - 'A') | r } // returns lower-case r if r is an ASCII letter
