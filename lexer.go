package algebra

import (
	"errors"
	"unicode"
)

const (
	NumberToken         = "Number"
	SkipToken           = "Skip"
	LeftBracketToken    = "LeftBracket"
	RightBracketToken   = "RightBracket"
	AdditionToken       = "Addition"
	SubtractionToken    = "Subtraction"
	MultiplicationToken = "Multiplication"
	DivisionToken       = "Division"
)

var (
	InvalidInputError = errors.New("lexer: invalid input")
)

type Token struct {
	Name     string
	Contents string
}

func (t Token) Type(n string) bool {
	if n == t.Name {
		return true
	}

	return false
}

type Lexer struct {
	source   string
	location int
}

func (lex *Lexer) Lex(source string) ([]Token, error) {
	lex.source = source
	lex.location = 0
	return lex.lex()
}

func (lex *Lexer) lex() ([]Token, error) {
	var stream []Token
	for !lex.End() {
		c := lex.source[lex.location]

		if lex.isNumber(c) {
			stream = append(stream, lex.number())
		} else if lex.isBrackets(c) {
			stream = append(stream, lex.brackets())
		} else if lex.isAddition(c) {
			stream = append(stream, lex.addition())
		} else if lex.isSubtraction(c) {
			stream = append(stream, lex.subtraction())
		} else if lex.isMultiplication(c) {
			stream = append(stream, lex.multiplication())
		} else if lex.isDivision(c) {
			stream = append(stream, lex.division())
		} else if lex.isWhitespace(c) {
			stream = append(stream, lex.whitespace())
		} else {
			return stream, InvalidInputError
		}

	}

	return stream, nil
}

func (lex *Lexer) addition() Token {
	lex.Next()

	return Token{AdditionToken, "+"}
}

func (lex *Lexer) subtraction() Token {
	lex.Next()

	return Token{SubtractionToken, "-"}
}

func (lex *Lexer) multiplication() Token {
	lex.Next()

	return Token{MultiplicationToken, "*"}
}

func (lex *Lexer) division() Token {
	lex.Next()

	return Token{DivisionToken, "/"}
}

func (lex *Lexer) whitespace() Token {
	var content string
	for !lex.End() && lex.isWhitespace(lex.source[lex.location]) {
		content += string(lex.source[lex.location])
		lex.Next()
	}

	return Token{SkipToken, content}
}

func (lex *Lexer) brackets() Token {
	c := lex.source[lex.location]

	lex.Next()
	if c == '(' {
		return Token{RightBracketToken, "("}
	}

	return Token{LeftBracketToken, ")"}
}

func (lex *Lexer) number() Token {
	var content string

	for !lex.End() && lex.isNumber(lex.source[lex.location]) {
		content += string(lex.source[lex.location])
		lex.Next()
	}

	return Token{NumberToken, content}
}

func (lex *Lexer) isSubtraction(c uint8) bool {
	if c == '-' {
		return true
	}

	return false
}

func (lex *Lexer) isAddition(c uint8) bool {
	if c == '+' {
		return true
	}

	return false
}

func (lex *Lexer) isMultiplication(c uint8) bool {
	if c == '*' {
		return true
	}

	return false
}

func (lex *Lexer) isDivision(c uint8) bool {
	if c == '/' {
		return true
	}

	return false
}

func (lex *Lexer) isWhitespace(c uint8) bool {
	return unicode.IsSpace(rune(c))
}

func (lex *Lexer) isBrackets(c uint8) bool {
	if c == '(' || c == ')' {
		return true
	}

	return false
}

func (lex *Lexer) isNumber(c uint8) bool {
	if c == '-' && unicode.IsNumber(rune(lex.source[lex.location+1])) {
		return true
	}

	if unicode.IsNumber(rune(c)) {
		return true
	}

	return false
}

func (lex *Lexer) End() bool {
	if lex.location >= len(lex.source) {
		return true
	}

	return false
}

func (lex *Lexer) Next() {
	lex.location += 1
}
