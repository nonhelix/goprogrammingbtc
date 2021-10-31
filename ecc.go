package main

import (
	"fmt"
	"math/big"
)

type FieldElement struct {
	num   *big.Int
	prime *big.Int
}

func Add(x, y FieldElement) FieldElement {
	if x.prime == y.prime {
		z := FieldElement{big.NewInt(int64(0)), x.prime}
		z.num.Add(x.num, y.num)
		z.num.Mod(z.num, x.prime)
		return z
	} else {
		//TODO: check how to return illegal fieldelement, not {-1,-1}
		fmt.Println("raise error, different prime")
		return FieldElement{nil, nil}
	}
}

func Sub(x, y FieldElement) FieldElement {
	if x.prime == y.prime {
		z := FieldElement{big.NewInt(int64(0)), x.prime}
		z.num.Sub(x.num, y.num)
		z.num.Mod(z.num, x.prime)
		return z
	} else {
		//TODO: check how to return illegal fieldelement, not {-1,-1}
		fmt.Println("raise error, different prime")
		return FieldElement{nil, nil}
	}
}

func Mul(x, y FieldElement) FieldElement {
	if x.prime == y.prime {
		z := FieldElement{big.NewInt(int64(0)), x.prime}
		z.num.Mul(x.num, y.num)
		z.num.Mod(z.num, x.prime)
		return z
	} else {
		//TODO: check how to return illegal fieldelement, not {-1,-1}
		fmt.Println("raise error, different prime")
		return FieldElement{nil, nil}
	}
}

func Pow(x FieldElement, n *big.Int) FieldElement {
	z := FieldElement{big.NewInt(int64(0)), x.prime}
	z.num.Exp(x.num, n, x.prime)
	return z
}

func Div(x, y FieldElement) FieldElement {
	if x.prime == y.prime {
		z := FieldElement{big.NewInt(int64(0)), x.prime}
		e := FieldElement{big.NewInt(int64(0)), x.prime}
		e.num.Sub(x.prime, big.NewInt(int64(2)))
		z.num.Exp(y.num, e.num, nil)
		z.num.Mul(z.num, x.num)
		z.num.Mod(z.num, x.prime)
		return z
	} else {
		//TODO: check how to return illegal fieldelement, not {-1,-1}
		fmt.Println("raise error, different prime")
		return FieldElement{nil, nil}
	}
}

type Point struct {
	a FieldElement
	b FieldElement
	x FieldElement
	y FieldElement
}

// y2 = x3  + 7 TODO make the validation.
//func (p PointFF) onCurveFF() bool {
//	return p.y.Pow(p.y, 2) == p.x.Pow(p.x, 3)
//}

func AddPoint(p1, p2 Point) (p3 Point) {
	//TODO check onCurve if both are Inf are also OK.(pg.34) and check sameCurve

	// same x and not same y case TODO check if p1.num > p1.prime is legal?
	InfFieldElement := FieldElement{nil, p1.x.prime}
	if p1.x == p2.x && p1.y != p2.y {
		return Point{p1.a, p1.b, InfFieldElement, InfFieldElement}
	}
	// Inf case TODO Inf is not big enough. Maybe all of them shouble be bytes?
	if p1.x.num == nil {
		return p2
	} else if p2.x.num == nil {
		return p1
	}
	if p1.x != p2.x {
		// s = (y2-y1)/(x2-x1)
		s := Div(Sub(p2.y, p1.y), Sub(p2.x, p1.x))
		fmt.Println("s", s.num)
		//x3 =s2 –x1 –x2
		x3 := Sub(Sub(Mul(s, s), p1.x), p2.x)
		//y3 = s(x1 – x3) – y1
		y3 := Sub(Mul(Sub(p1.x, x3), s), p1.y)
		p3 = Point{a: p3.a, b: p3.b, x: x3, y: y3}
		return p3
	}

	// P1 == P2 and y = 0
	if p1 == p2 && p1.y.num == nil {
		return Point{p1.a, p1.b, InfFieldElement, InfFieldElement}
	}
	// P1 == P2
	if p1 == p2 {
		// s = (3x12 + a)/(2y1)
		s := Div(Add(Add(Mul(p1.x, p1.x), Mul(p1.x, p1.x)), Mul(p1.x, p1.x)), Add(p1.y, p1.y))

		// x3 = s2 – 2x1
		x3 := Sub(Mul(s, s), Add(p1.x, p1.x))

		// y3 = s(x1 – x3) – y1

		y3 := Sub(Mul(Sub(p1.x, x3), s), p1.y)
		p3 = Point{a: p3.a, b: p3.b, x: x3, y: y3}
		return p3
	}
	return
}
func main() {
	big0 := big.NewInt(int64(0))
	big223 := big.NewInt(int64(223))
	big7 := big.NewInt(int64(7))
	big47 := big.NewInt(int64(47))
	big71 := big.NewInt(int64(71))
	big17 := big.NewInt(int64(17))
	big56 := big.NewInt(int64(56))

	big143 := big.NewInt(int64(143))
	big98 := big.NewInt(int64(98))
	big76 := big.NewInt(int64(76))
	big66 := big.NewInt(int64(66))
	a := FieldElement{big0, big223}
	b := FieldElement{big7, big223}
	x1 := FieldElement{big47, big223}
	y1 := FieldElement{big71, big223}
	x2 := FieldElement{big17, big223}
	y2 := FieldElement{big56, big223}
	p1 := Point{a, b, x1, y1}
	p2 := Point{a, b, x2, y2}
	x4 := FieldElement{big143, big223}
	y4 := FieldElement{big98, big223}
	x5 := FieldElement{big76, big223}
	y5 := FieldElement{big66, big223}
	p4 := Point{a, b, x4, y4}
	p5 := Point{a, b, x5, y5}
	p3 := AddPoint(p1, p2)
	p6 := AddPoint(p4, p5)
	fmt.Println("p3.x.num:", p3.x.num)
	fmt.Println("p3.y.num:", p3.y.num)
	fmt.Println("p6.x.num:", p6.x.num)
	fmt.Println("p6.y.num:", p6.y.num)
}
