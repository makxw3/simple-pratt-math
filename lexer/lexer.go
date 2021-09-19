package lexer

import (
	"pratt/token"
)

// The lexer struct
type Lexer struct {
	input                  string // The input to be scanned
	currentReadingPosition int    // The index of the current reading position
	peekReadingPosition    int    // The index of the next reading position.
	char                   byte   // The current byte being read
}

// A helper function to create a new Lexer object
func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	// Read the next char
	lexer.readNextChar()
	return lexer
}

func (lx *Lexer) readNextChar() {
	if lx.peekReadingPosition >= len(lx.input) {
		// Return the null byte for token.EOF
		lx.char = 0
	} else {
		// Read the next byte from the input and set it in lx.char
		lx.char = lx.input[lx.peekReadingPosition]
	}
	// Advance the pointers
	lx.currentReadingPosition = lx.peekReadingPosition
	lx.peekReadingPosition++
}

func (lx *Lexer) ReadNextToken() token.Token {
	// Skip all the white spaces until lx.char is not a white space
	lx.skipWhiteSpaces()
	var tok token.Token
	switch lx.char {
	case '+':
		tok = makeToken(token.ADD, lx.char)
	case '-':
		tok = makeToken(token.MINUS, lx.char)
	case '*':
		tok = makeToken(token.ASTERISK, lx.char)
	case 0:
		tok = makeToken(token.EOF, 0)
	default:
		if isDigit(lx.char) {
			return lx.scanNumber()
		}
		tok = token.Token{Type: token.ILLEGAL, Literal: "'" + string(lx.char) + "'"}
	}
	lx.readNextChar()
	return tok
}

// scanNumber scans the number until it encounters a non-digit
func (lx *Lexer) scanNumber() token.Token {
	currentPosition := lx.currentReadingPosition
	for isDigit(lx.char) {
		lx.readNextChar()
	}
	value := lx.input[currentPosition:lx.currentReadingPosition]
	return token.Token{Type: token.NUMBER, Literal: value}
}

// Helper function to check if a byte is a number
func isDigit(char byte) bool {
	return char <= '9' && char >= '0'
}

// Helper function to create new tokens
func makeToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

// Helper function to skip white spaces
func (lx *Lexer) skipWhiteSpaces() {
	for lx.char == ' ' || lx.char == '\t' || lx.char == '\r' || lx.char == '\n' {
		lx.readNextChar()
	}
}
