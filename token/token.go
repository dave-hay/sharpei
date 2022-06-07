package token

type TType string

type Token struct {
	Type    TType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // unknown token
	EOF     = "EOF"     // end of file

	// Identifiers + literals
	IDENT = "IDENt" // add, foobar, x, y
	INT   = "Int"   //123

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	// Conditionals
	EQ     = "=="
	NOT_EQ = "!="
)

var keywords = map[string]TType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checks the keyword table to see if ident is a keyword
func LookupIdent(ident string) TType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
