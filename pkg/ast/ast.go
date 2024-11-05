package ast

// These are the available root node types. In JSON it will either be an
// object or an array at the base.
const (
	ObjectRoot RootNodeType = iota
	ArrayRoot
)

// RootNodeType is a type alias for an int
type RootNodeType int

// RootNode is what starts every parsed AST. There is a `Type` field so that
// you can ask which root node type starts the tree.
type RootNode struct {
	RootValue *Value
	Type      RootNodeType
}

// Available ast value types
const (
	ObjectType Type = iota
	ArrayType
	ArrayItemType
	LiteralType
	PropertyType
	IdentifierType
)

const (
	StringLiteralValueType LiteralValueType = iota
	NumberLiteralValueType
	NullLiteralValueType
	BooleanLiteralValueType
)
