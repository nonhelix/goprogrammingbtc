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
