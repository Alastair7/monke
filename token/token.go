package token

const (
	ILLEGAL = "ILLEGAL" // Unknown Token will be Illegal.
	EOF     = "EOF"

	// Identifier Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}