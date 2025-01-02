package parser

import "testing"

func TestTransFormExp(t *testing.T) {
	s := "header=\"realm=\\\"Comtrend Gigabit 802.11n Router\" || banner=\"Comtrend Gigabit 802.11n Router\""
	tokens, err := ParseTokens(s)
	if err != nil {
		t.Fatal(err)
	}
	exp, err := TransFormExp(tokens)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(exp)
}
