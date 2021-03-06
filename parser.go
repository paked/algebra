package algebra

import (
	"errors"
	"strconv"
)

func Evaluate(source string) (int, error) {
	l := Lexer{}
	tokens, err := l.Lex(source)
	if err != nil {
		return -1, err
	}

	p := Parser{}
	n, err := p.Parse(tokens)
	if err != nil {
		return -1, err
	}
	return n.Eval(), nil
}

type Parser struct {
	tokens   []Token
	location int
}

func (p *Parser) addition() (Node, error) {
	node, err := p.subtraction()
	if err != nil {
		return node, err
	}

	t, err := p.Peek()
	if err != nil {
		return node, err
	}

	if !t.Type(AdditionToken) {
		return node, nil
	}

	p.Next()
	right, err := p.addition()
	if err != nil {
		return node, err
	}
	node = AdditionNode{Left: node, Right: right}

	return node, nil
}

func (p *Parser) subtraction() (Node, error) {
	node, err := p.multiplication()
	if err != nil {
		return node, err
	}

	t, err := p.Peek()
	if err != nil {
		return node, err
	}

	if !t.Type(SubtractionToken) {
		return node, nil
	}

	p.Next()
	right, err := p.subtraction()
	if err != nil {
		return node, err
	}
	node = SubtractionNode{Left: node, Right: right}

	return node, nil
}

func (p *Parser) multiplication() (Node, error) {
	node, err := p.division()
	if err != nil {
		return node, err
	}

	t, err := p.Peek()
	if err != nil {
		return node, err
	}

	if !t.Type(MultiplicationToken) {
		return node, nil
	}

	p.Next()
	right, err := p.multiplication()
	if err != nil {
		return node, err
	}
	node = MultiplicationNode{Left: node, Right: right}

	return node, nil
}

func (p *Parser) division() (Node, error) {
	node, err := p.exponents()
	if err != nil {
		return node, err
	}

	t, err := p.Peek()
	if err != nil {
		return node, err
	}

	if !t.Type(DivisionToken) {
		return node, nil
	}

	p.Next()
	right, err := p.division()
	if err != nil {
		return node, err
	}
	node = DivisionNode{Left: node, Right: right}

	return node, nil
}

func (p *Parser) exponents() (Node, error) {
	node, err := p.expression()
	if err != nil {
		return node, err
	}

	t, err := p.Peek()
	if err != nil {
		return node, err
	}

	if !t.Type(PowerToken) {
		return node, nil
	}

	p.Next()
	right, err := p.exponents()
	if err != nil {
		return node, err
	}
	node = PowerNode{Left: node, Right: right}

	return node, err
}

func (p *Parser) expression() (Node, error) {
	t, err := p.Peek()
	if !t.Type(RightBracketToken) {
		if !t.Type(NumberToken) {
			return NumberNode{}, errors.New("Invalid token!")
		}

		i, err := strconv.Atoi(t.Contents)
		if err != nil {
			return NumberNode{}, err
		}
		node := NumberNode{i}
		p.Next()
		return node, nil
	}

	p.Next()
	node, err := p.addition()
	if err != nil {
		return node, err
	}

	if t, err := p.Peek(); err == nil && !t.Type(LeftBracketToken) {
		return node, errors.New("Expecting left bracket, didnt get it!")
	}

	p.Next()
	return node, nil
}

func (p *Parser) Parse(tokens []Token) (Node, error) {
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
		return Token{}, nil
	}

	return p.tokens[p.location], nil
}

func (p *Parser) Next() {
	p.location += 1
}
