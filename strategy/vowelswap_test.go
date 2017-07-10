package strategy_test

import (
	"testing"

	"../strategy"
)

func TestVowelSwap(t *testing.T) {
	out, err := strategy.VowelSwap.Generate("zenithar")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 18 {
		t.FailNow()
	}
}
