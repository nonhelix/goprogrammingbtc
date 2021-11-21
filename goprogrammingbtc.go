package goprogrammingbtc

import (
	"math/big"
)

type FieldElement struct {
	num   *big.Int
	prime *big.Int
}

func (x FieldElement) Add(y FieldElement) FieldElement {
	if x.prime.Cmp(y.prime) != 0 {
		panic("mismatched primes in fieldElement")
	}
	z := FieldElement{big.NewInt(0), x.prime}
	z.num.Add(x.num, y.num)
	z.num.Mod(z.num, x.prime)
	return z
}
func (x FieldElement) Sub(y FieldElement) FieldElement {
	if x.prime.Cmp(y.prime) != 0 {
		panic("mismatched primes in fieldElement")
	}
	z := FieldElement{big.NewInt(0), x.prime}
	z.num.Sub(x.num, y.num)
	z.num.Mod(z.num, x.prime)
	return z
}
func (x FieldElement) Mul(y FieldElement) FieldElement {
	if x.prime.Cmp(y.prime) != 0 {
		panic("mismatched primes in fieldElement")
	}
	z := FieldElement{big.NewInt(0), x.prime}
	z.num.Mul(x.num, y.num)
	z.num.Mod(z.num, x.prime)
	return z
}
func (x FieldElement) Pow(n *big.Int) FieldElement {
	z := FieldElement{big.NewInt(0), x.prime}
	// make it negative compatiable
	n.Mod(n, big.NewInt(0).Sub(x.prime, big.NewInt(1)))
	z.num.Exp(x.num, n, x.prime)
	return z
}
func (x FieldElement) Div(y FieldElement) FieldElement {
	if x.prime.Cmp(y.prime) != 0 {
		panic("mismatched primes in fieldElement")
	}
	z := FieldElement{big.NewInt(0), x.prime}
	z.num.Exp(y.num, big.NewInt(0).Sub(x.prime, big.NewInt(int64(2))), x.prime)
	z.num.Mul(z.num, x.num)
	z.num.Mod(z.num, x.prime)
	return z
}
