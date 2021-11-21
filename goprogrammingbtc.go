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
	z := FieldElement{big.NewInt(int64(0)), x.prime}
	z.num.Add(x.num, y.num)
	z.num.Mod(z.num, x.prime)
	return z
}
func (x FieldElement) Sub(y FieldElement) FieldElement {
	if x.prime.Cmp(y.prime) != 0 {
		panic("mismatched primes in fieldElement")
	}
	z := FieldElement{big.NewInt(int64(0)), x.prime}
	z.num.Sub(x.num, y.num)
	z.num.Mod(z.num, x.prime)
	return z
}
func (x FieldElement) Mul(y FieldElement) FieldElement {
	if x.prime.Cmp(y.prime) != 0 {
		panic("mismatched primes in fieldElement")
	}
	z := FieldElement{big.NewInt(int64(0)), x.prime}
	z.num.Mul(x.num, y.num)
	z.num.Mod(z.num, x.prime)
	return z
}
func (x FieldElement) Pow(n *big.Int) FieldElement {
	z := FieldElement{big.NewInt(int64(0)), x.prime}
	z.num.Exp(x.num, n, x.prime)
	return z
}
