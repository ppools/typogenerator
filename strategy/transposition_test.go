package strategy_test

import (
	"testing"

	"../strategy"
)

func TestTransposition(t *testing.T) {
	out, err := strategy.Transposition.Generate("zemithar")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 7 {
		t.FailNow()
	}
}
