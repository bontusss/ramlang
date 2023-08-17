package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Illegal chars and unknown identifiers or literals
	EOF = "EOF" // End of file: tells the parser when to stop.

	// Identifiers + Literals
	IDENT = "IDENT" // identfiers: add, foo, x, y ...
	INT = "INT" // Intergers
	 
	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiter
	COMMA = ","
	SEMICOLON = ";"
	LEFT_PARA = "("
	RIGHT_PARA = ")"
	LEFTBRACE = "{"
	RIGHTBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"

)