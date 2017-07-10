package strategy_test

import (
	"testing"

	"../mapping"
	"../strategy"
)

func TestDoubleHit(t *testing.T) {
	out, err := strategy.DoubleHit(mapping.French).Generate("zenithar")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 77 {
		t.FailNow()
	}
}
