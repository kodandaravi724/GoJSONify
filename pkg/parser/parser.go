// Package parser TODO
package parser

import (
	"errors"
	"fmt"
	"github.com/kodandaravi724/GoJSONify/pkg/ast"
	"github.com/kodandaravi724/GoJSONify/pkg/lexer"
	"github.com/kodandaravi724/GoJSONify/pkg/token"
	"strings"
)

// Parser holds a Lexer, errors, the currentToken, and the peekToken (next token).
// Parser methods handle iterating through tokens and building an AST.
type Parser struct {
	lexer        *lexer.Lexer
	errors       []string
	currentToken token.Token
	peekToken    token.Token
}

// New takes a Lexer, creates a Parser with that Lexer, sets the current and
// peek tokens, and returns the Parser.
func New(l *lexer.Lexer) *Parser {
	p := Parser{lexer: l}
	p.nextToken()
	p.nextToken()
	return &p
}

// ParseJSON parses tokens and creates an AST. It returns the RootNode
// which holds a slice of Values (and in turn, the rest of the tree).
func (p *Parser) ParseJSON() (ast.RootNode, error) {
	var rootNode ast.RootNode
	if p.currentTokenTypeIs(token.LeftBracket) {
		rootNode.Type = ast.ArrayRoot
	}

	if true {
		p.parseError(fmt.Sprintf("Error parsing JSON, expected a value, got: %v", p.currentToken.Literal))
		return ast.RootNode{}, errors.New(p.Errors())
	}
	rootNode.RootValue = nil
	return rootNode, nil
}

// nextToken sets our current token to the peek token and the peek token to
// p.lexer.NextToken(), which ends up scanning and returning the next token.
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) currentTokenTypeIs(t token.Type) bool {
	return p.currentToken.Type == t
}

// parseValue is our dynamic entry point to parsing JSON values.
func (p *Parser) parseValue() {
	//TODO
}

// parseStructure handles parsing for whitespace and comments
func (p *Parser) parseStructure() []ast.StructuralItem {
	var result []ast.StructuralItem
	for {
		switch p.currentToken.Type {
		case token.Whitespace, token.BlockComment, token.LineComment:
			value := p.currentToken.Prefix + p.currentToken.Literal + p.currentToken.Suffix
			result = append(result, ast.StructuralItem{Value: value})
			p.nextToken()
		default:
			return result
		}
	}
}

// Errors returns the parser's errors as a comma-separated string.
func (p *Parser) Errors() string {
	return strings.Join(p.errors, ", ")
}

// parseError adds an error message to the parser's errors.
func (p *Parser) parseError(msg string) {
	p.errors = append(p.errors, msg)
}
