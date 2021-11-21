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
func TestExercise1_4(t *testing.T) {
	prime := big.NewInt(int64(97))
	x1 := FieldElement{big.NewInt(int64(95)), prime}
	y1 := FieldElement{big.NewInt(int64(45)), prime}
	z1 := FieldElement{big.NewInt(int64(31)), prime}
	res1 := x1.Mul(y1).Mul(z1)
	ans1 := big.NewInt(int64(23))
	if res1.num.Cmp(ans1) != 0 {
		t.Error("wrong res1", res1)
	}
	x2 := FieldElement{big.NewInt(int64(17)), prime}
	y2 := FieldElement{big.NewInt(int64(13)), prime}
	z2 := FieldElement{big.NewInt(int64(19)), prime}
	a2 := FieldElement{big.NewInt(int64(44)), prime}
	res2 := x2.Mul(y2).Mul(z2).Mul(a2)
	ans2 := big.NewInt(int64(68))
	if res2.num.Cmp(ans2) != 0 {
		t.Error("wrong res2", res2)
	}
	x3 := FieldElement{big.NewInt(int64(12)), prime}
	y3 := big.NewInt(int64(7))
	z3 := FieldElement{big.NewInt(int64(77)), prime}
	a3 := big.NewInt(int64(49))
	one := x3.Pow(y3)
	two := z3.Pow(a3)
	res3 := one.Mul(two)
	ans3 := big.NewInt(int64(63))
	if res3.num.Cmp(ans3) != 0 {
		t.Error("wrong res3", res3)
	}
}
func TestExercise1_8(t *testing.T) {
	prime := big.NewInt(int64(31))
	x1 := FieldElement{big.NewInt(int64(3)), prime}
	y1 := FieldElement{big.NewInt(int64(24)), prime}
	res1 := x1.Div(y1)
	ans1 := big.NewInt(int64(4))
	if res1.num.Cmp(ans1) != 0 {
		t.Error("wrong res1", res1)
	}
	x2 := FieldElement{big.NewInt(int64(17)), prime}
	y2 := FieldElement{big.NewInt(int64(-3)), prime}
	res2 := x2.Pow(y2.num)
	ans2 := big.NewInt(int64(29))
	if res2.num.Cmp(ans2) != 0 {
		t.Error("wrong res2", res2)
	}
	x4 := FieldElement{big.NewInt(int64(4)), prime}
	y4 := FieldElement{big.NewInt(int64(-4)), prime}
	z4 := FieldElement{big.NewInt(int64(11)), prime}
	res4 := x4.Pow(y4.num).Mul(z4)
	ans4 := big.NewInt(int64(13))
	if res4.num.Cmp(ans4) != 0 {
		t.Error("wrong res4", res4)
	}
}
