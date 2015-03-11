package algebra

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	l := Lexer{"1 + 90 hello", 0}
	fmt.Println(l.lex())
}
