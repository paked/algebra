package algebra

import (
	"errors"
	"fmt"
	"strconv"
)

func Evaluate(source string) int {
	l := Lexer{source, 0}
	tokens := l.Lex()
	fmt.Println("tokens are:", tokens)

	p := Parser{}
	n := p.Parse(tokens)

	return n.Eval()
}

type Parser struct {
	tokens   []Token
	location int
}

func (p *Parser) addition() Node {
	node := p.subtraction()
	if t, err := p.Peek(); err == nil && t.Type(AdditionToken) {
		p.Next()
		right := p.addition()
		node = AdditionNode{Left: node, Right: right}
	}

	return node
}

func (p *Parser) subtraction() Node {
	node := p.multiplication()

	if t, err := p.Peek(); err == nil && t.Type(SubtractionToken) {
		p.Next()
		right := p.subtraction()
		node = SubtractionNode{Left: node, Right: right}
	}

	return node
}

func (p *Parser) multiplication() Node {
	node := p.division()
	if t, err := p.Peek(); err == nil && t.Type(MultiplicationToken) {
		p.Next()
		right := p.multiplication()
		node = MultiplicationNode{Left: node, Right: right}
	}

	return node
}

func (p *Parser) division() Node {
	node := p.expression()
	if t, err := p.Peek(); err == nil && t.Type(DivisionToken) {
		p.Next()
		right := p.division()
		node = DivisionNode{Left: node, Right: right}
	}

	return node
}

func (p *Parser) expression() Node {
	if t, err := p.Peek(); err == nil && t.Type(RightBracketToken) {
		p.Next()
		node := p.addition()
		if t, err := p.Peek(); err == nil && !t.Type(LeftBracketToken) {
			fmt.Println("No left bracket!")
			return nil
		}
		p.Next()
		return node
	}

	if t, err := p.Peek(); err == nil && t.Type(NumberToken) {
		i, _ := strconv.Atoi(t.Contents)
		node := NumberNode{i}
		p.Next()
		fmt.Println("created number node!")
		return node
	} else {
		fmt.Println("Error:", err, "Token:", t)
	}

	return nil
}

func (p *Parser) Parse(tokens []Token) Node {
	p.tokens = p.cleanseInput(tokens)

	return p.addition()
}

func (p *Parser) cleanseInput(tokens []Token) []Token {
	var output []Token

	for _, token := range tokens {
		if !token.Type(SkipToken) {
			output = append(output, token)
		}
	}

	fmt.Println("cleansed tokens are:", output)

	return output
}

func (p *Parser) End() bool {
	if p.location >= len(p.tokens) {
		return true
	}

	return false
}

func (p *Parser) Peek() (Token, error) {
	if p.End() {
		fmt.Println("no more tokens..")
		return Token{}, errors.New("End of tokens...")
	}

	return p.tokens[p.location], nil
}

func (p *Parser) Next() {
	p.location += 1
}
