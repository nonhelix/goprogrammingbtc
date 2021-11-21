package goprogrammingbtc

import (
	"math/big"
	"testing"
)

func TestExercise1_2(t *testing.T) {
	prime := big.NewInt(int64(57))
	x1 := FieldElement{big.NewInt(int64(44)), prime}
	y1 := FieldElement{big.NewInt(int64(33)), prime}
	res1 := x1.Add(y1)
	ans1 := big.NewInt(int64(20))
	if res1.num.Cmp(ans1) != 0 {
		t.Error("wrong res1", res1)
	}
}
