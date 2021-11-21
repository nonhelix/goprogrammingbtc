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
	x2 := FieldElement{big.NewInt(int64(9)), prime}
	y2 := FieldElement{big.NewInt(int64(29)), prime}
	res2 := x2.Sub(y2)
	ans2 := big.NewInt(int64(37))
	if res2.num.Cmp(ans2) != 0 {
		t.Error("wrong res2", res2)
	}
	x3 := FieldElement{big.NewInt(int64(17)), prime}
	y3 := FieldElement{big.NewInt(int64(42)), prime}
	z3 := FieldElement{big.NewInt(int64(49)), prime}
	res3 := x3.Add(y3).Add(z3)
	ans3 := big.NewInt(int64(51))
	if res3.num.Cmp(ans3) != 0 {
		t.Error("wrong res3", res3)
	}
	x4 := FieldElement{big.NewInt(int64(52)), prime}
	y4 := FieldElement{big.NewInt(int64(30)), prime}
	z4 := FieldElement{big.NewInt(int64(38)), prime}
	res4 := x4.Sub(y4).Sub(z4)
	ans4 := big.NewInt(int64(41))
	if res4.num.Cmp(ans4) != 0 {
		t.Error("wrong res4", res4)
	}
}
