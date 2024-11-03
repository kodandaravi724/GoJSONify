package token

import (
	"fmt"
)

// All the different tokens for supporting JSON
const (
	// Token/character we don't know about
	Illegal Type = "ILLEGAL"

	// End of file
	EOF Type = "EOF"

	// Literals
	String Type = "STRING"
	Number Type = "NUMBER"

	// The six structural tokens
	LeftBrace    Type = "{"
	RightBrace   Type = "}"
	LeftBracket  Type = "["
	RightBracket Type = "]"
	Comma        Type = ","
	Colon        Type = ":"

	// Structural
	Whitespace Type = "WHITESPACE"

	// Comments
	LineComment  Type = "//"
	BlockComment Type = "/*"

	// Values
	True  Type = "TRUE"
	False Type = "FALSE"
	Null  Type = "NULL"
)

// Type is a type alias for a string
type Type string

// Token is a struct representing a JSON token - It holds information like its Type and Literal, as well
// as Start, End, and Line fields. Line is used for better error handling, while Start and End are used
// to return objects/arrays from querys.
type Token struct {
	Type    Type
	Literal string
	Line    int
	Start   int
	End     int
	Prefix  string
	Suffix  string
	Reason  string // optional reason when Illegal Type
}

var validJSONIdentifiers = map[string]Type{
	"true":  True,
	"false": False,
	"null":  Null,
}

// LookupIdentifier checks our validJSONIdentifiers map for the scanned identifier. If it finds one,
// the identifier's token type is returned. If not found, an error is returned
func LookupIdentifier(identifier string) (Type, error) {
	if token, ok := validJSONIdentifiers[identifier]; ok {
		return token, nil
	}
	return "", fmt.Errorf("Expected a valid JSON identifier. Found: %s", identifier)
}

var escapes = map[rune]int{
	'"':  0, // Quotation mask
	'\\': 1, // Reverse solidus
	'/':  2, // Solidus
	'b':  3, // Backspace
	'f':  4, // Form feed
	'n':  5, // New line
	'r':  6, // Carriage return
	't':  7, // Horizontal tab
	'u':  8, // 4 hexadecimal digits
}

var escapeChars = map[string]string{
	"b": "\b", // Backspace
	"f": "\f", // Form feed
	"n": "\n", // New line
	"r": "\r", // Carriage return
	"t": "\t", // Horizontal tab
}
