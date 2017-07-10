package strategy_test

import (
	"testing"

	"../strategy"
)

func TestAddition(t *testing.T) {
	out, err := strategy.Addition.Generate("zenithar")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 26 {
		t.FailNow()
	}
}
