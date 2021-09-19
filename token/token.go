package token

import "fmt"

// The type of the token
type TokenType string

// The list of token types
const (
	EOF     = "EOF"     // The End-Of-File token
	ILLEGAL = "ILLEGAL" // An illegal token
	NUMBER  = "NUMBER"  // The number token eg 10, 20, 40

	// Operators
	ADD      = "+" // The `+` operator in `1 + 2`
	MINUS    = "-" // The `-` operator in `10 - 1`
	ASTERISK = "*" // The `*` operator in `23 * 10`
)

// The Token struct
type Token struct {
	Type    TokenType // The token-type eg EOF
	Literal string    // The literal value of the token eg age for the token <IDENT, age>
}

// Helper function to print the token
func (tk *Token) Print() {
	fmt.Printf("<%v,%v>\n", tk.Type, tk.Literal)
}
