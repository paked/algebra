package algebra

import (
	"fmt"
	"testing"
)

var l Lexer

func TestTokens(t *testing.T) {
	fmt.Println(l.Lex("1 + 90 (1 + 2 ) - 3"))
}
