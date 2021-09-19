package parser

import (
	"fmt"
	"pratt/ast"
	"pratt/lexer"
	"pratt/token"
	"strconv"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token // The current token being read
	peekToken    token.Token // The next token to be read

	// Pratt Parsing Functions
	prefixParsingFunctins map[token.TokenType]prefixParsingFunction
	infixParsingFunctions map[token.TokenType]infixParsingFunction
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer}
	// Set the currentToken and the peekToken
	parser.advance()
	parser.advance()

	// Initialize the maps
	parser.prefixParsingFunctins = make(map[token.TokenType]prefixParsingFunction)
	parser.infixParsingFunctions = make(map[token.TokenType]infixParsingFunction)

	// Add the parsing functions
	parser.addPrefixParsingFunction(token.NUMBER, parser.parseNumberExpression)

	// Add the infix parsing functions
	parser.addInfixParsingFunction(token.ADD, parser.parseInfixExpression)
	parser.addInfixParsingFunction(token.MINUS, parser.parseInfixExpression)
	parser.addInfixParsingFunction(token.ASTERISK, parser.parseInfixExpression)
	return parser
}

// Advanced the tokens
func (ps *Parser) advance() {
	ps.currentToken = ps.peekToken
	ps.peekToken = ps.lexer.ReadNextToken()
}

// The types of parsing functions
type (
	prefixParsingFunction func() ast.Expression
	infixParsingFunction  func(ast.Expression) ast.Expression
)

// Helper function for adding prefixParsingFunctions
func (ps *Parser) addPrefixParsingFunction(tokenType token.TokenType, fn prefixParsingFunction) {
	ps.prefixParsingFunctins[tokenType] = fn
}

// Helper functions for adding infixParsingFunctions
func (ps *Parser) addInfixParsingFunction(tokenType token.TokenType, fn infixParsingFunction) {
	ps.infixParsingFunctions[tokenType] = fn
}

/** The binding power **/
var OperatorBindingPower = map[token.TokenType]int{
	token.ASTERISK: 30, // The highest binding power
	token.ADD:      10,
	token.MINUS:    10,
}

// The lowest binding power
const LOWEST = 0

// Helper function to get the binding power of the current token
func (ps *Parser) currentBp() int {
	bp, ok := OperatorBindingPower[ps.currentToken.Type]
	if !ok {
		return LOWEST
	}
	return bp
}

// Helper function to get the binding power of the next token
func (ps *Parser) peekBp() int {
	bp, ok := OperatorBindingPower[ps.peekToken.Type]
	if !ok {
		return LOWEST
	}
	return bp
}

/** Parsing Functions **/
func (ps *Parser) parseNumberExpression() ast.Expression {
	res, _ := strconv.ParseInt(ps.currentToken.Literal, 0, 0)
	return &ast.SingleNumberExpression{Value: res}
}

func (ps *Parser) parseInfixExpression(leftExpresion ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Operator: ps.currentToken,
		Left:     leftExpresion,
	}
	// Get the binding power of the current operator
	currentBp := ps.currentBp()
	// Advance to the next operator
	ps.advance()
	// Parse the right expression
	expression.Right = ps.parseExpression(currentBp)
	return expression
}

func (ps *Parser) parseExpression(bindingPower int) ast.Expression {
	// Look for a prefix parsing function for the current token
	prefixParsinFn, ok := ps.prefixParsingFunctins[ps.currentToken.Type]
	if !ok {
		fmt.Printf("Error! no prefixParsingFunction found for %v\n", ps.currentToken)
	}
	// parse the expression and assign the result to `leftExpression`
	leftExpression := prefixParsinFn()

	// This is where you handle the operator precedence and associativity
	if bindingPower == ps.peekBp() {
		bindingPower--
	}

	for bindingPower < ps.peekBp() {
		infixParsingFn := ps.infixParsingFunctions[ps.peekToken.Type]
		ps.advance()
		leftExpression = infixParsingFn(leftExpression)
	}
	return leftExpression
}

func (ps *Parser) ParseMathExpression() ast.Expression {
	return ps.parseExpression(LOWEST)
}
