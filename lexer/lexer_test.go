package lexer

import (
	"testing"

	"github.com/bontusss/ramlang/token"
)

func TestNextToken(t *testing.T) {
	input := `+=(){,;}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LEFT_PARA, "("},
		{token.RIGHT_PARA, ")"},
		{token.RIGHTBRACE, "}"},	
		{token.LEFTBRACE, "{"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for key, value := range tests {
		tok := l.NextToken()

		if tok.Type != value.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong, expected %q, got %q", key, value.expectedType, tok.Type)
		}

		if tok.Literal != value.expectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong, expected %q, got %q", key, value.expectedLiteral, tok.Literal)
		}
	}
}
