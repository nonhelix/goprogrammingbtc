package main

import (
	"fmt"
	"math/big"
)

type FieldElement struct {
	num   *big.Int
	prime *big.Int
}

func (z FieldElement) Add(x, y FieldElement) FieldElement {
	if x.prime == y.prime {
		z.num.Mod(x.num.Add(x.num, y.num), x.prime)
		z.prime = x.prime
		return z
	} else {
		//TODO: check how to return illegal fieldelement, not {-1,-1}
		fmt.Println("raise error, different prime")
		return z
	}
}

func (z FieldElement) Sub(x, y FieldElement) FieldElement {
	if x.prime == y.prime {
		z.num.Mod(x.num.Sub(x.num, y.num), x.prime)
		z.prime = x.prime
		return z
	} else {
		//TODO: check how to return illegal fieldelement, not {-1,-1}
		fmt.Println("raise error, different prime")
		return z
	}
}

func (z FieldElement) Mul(x, y FieldElement) FieldElement {
	if x.prime == y.prime {
		z.num.Mod(x.num.Mul(x.num, y.num), x.prime)
		z.prime = x.prime
		return z
	} else {
		//TODO: check how to return illegal fieldelement, not {-1,-1}
		fmt.Println("raise error, different prime")
		return z
	}
}

func (z FieldElement) Pow(x FieldElement, n *big.Int) FieldElement {
	z.num.Exp(x.num, n, x.prime)
	z.prime = x.prime
	return z
}

func (z FieldElement) Div(x, y FieldElement) FieldElement {
	if x.prime == y.prime {
		z.num.Exp(y.num, y.num.Sub(x.prime, big.NewInt(int64(2))), nil).Mul(z.num, x.num).Mod(z.num, x.prime)
		z.prime = x.prime
		return z
	} else {
		//TODO: check how to return illegal fieldelement, not {-1,-1}
		fmt.Println("raise error, different prime")
		return z
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

func (p1 Point) AddPoint(p2 Point) (p3 Point) {
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
		// initialize p3
		p3 = p1
		fmt.Println("p3.x.num", p3.x.num)

		fmt.Println("p3.y.num", p3.y.num)
		fmt.Println("p3.a.num", p3.a.num)
		//TODO  WARNING: why b is 54??? its pointer???
		fmt.Println("p3.b.num", p3.b.num)
		// initialize s. as a fieldelement, s can only be initialized as exited fieldelement.
		s := p1.x
		s.Sub(p2.y, p1.y)
		x2subx1 := p1.x
		x2subx1.Sub(p2.x, p1.x)
		s.Div(s, x2subx1)
		//fmt.Println("s======", s.num, s.prime)
		//x3 =s2 –x1 –x2
		///	var x3 FieldElement
		p3.x.Mul(s, s).Sub(p3.x, p1.x).Sub(p3.x, p2.x)
		fmt.Println("p3.x.num=", p3.x.num)
		//y3 = s(x1 – x3) – y1
		///	var y3 FieldElement
		p3.x = FieldElement{p3.x.num, p3.x.prime}
		fmt.Println("p3.x.num ============", p3.x.num)
		p3.y.Sub(p1.x, p3.x).Mul(p3.y, s).Sub(p3.y, p1.y)
		p3.y = FieldElement{p3.y.num, p3.y.prime}
		p3 = Point{a: p3.a, b: p3.b, x: p3.x, y: p3.y}
		fmt.Println("p3.x.num==================", p3.x.num)
		return p3
	}

	// P1 == P2 and y = 0
	if p1 == p2 && p1.y.num == nil {
		return Point{p1.a, p1.b, InfFieldElement, InfFieldElement}
	}
	// P1 == P2
	if p1 == p2 {
		// s = (3x12 + a)/(2y1)
		s := p1.x
		// initiazed 2y1
		twoy1 := p1.x
		twoy1.Add(p1.y, p1.y)
		s.Mul(p1.x, p1.x).Add(s, s).Add(s, s).Add(s, p1.a).Div(s, twoy1)

		// x3 = s2 – 2x1
		x3 := p1.x
		twox1 := p1.x
		twox1.Add(p1.x, p1.x)
		x3.Mul(s, s).Sub(x3, twox1)

		// y3 = s(x1 – x3) – y1

		y3 := p1.x
		y3.Sub(p1.x, x3).Mul(y3, s).Sub(y3, p1.y)
		return Point{p1.a, p1.b, x3, y3}
	}
	return
}

func (p1 Point) AddPointtest(p2 Point) (p3 Point) {

	// s = (y2-y1)/(x2-x1)
	p3 = p1
	// initialize s. as a fieldelement, s can only be initialized as exited fieldelement.
	s := p1.x
	s.Sub(p2.y, p1.y)
	x2subx1 := p1.x
	x2subx1.Sub(p2.x, p1.x)
	s.Div(s, x2subx1)
	//fmt.Println(s)
	//x3 =s2 –x1 –x2
	///	var x3 FieldElement
	p3.x.Mul(s, s).Sub(p3.x, p1.x).Sub(p3.x, p2.x)
	//y3 = s(x1 – x3) – y1
	///	var y3 FieldElement

	p3.y.Sub(p1.x, p3.x).Mul(p3.y, s).Sub(p3.y, p1.y)
	return Point{p3.a, p3.b, p3.x, p3.y}
}
func main() {
	big0 := big.NewInt(int64(0))
	big223 := big.NewInt(int64(223))
	big7 := big.NewInt(int64(7))
	big47 := big.NewInt(int64(47))
	big71 := big.NewInt(int64(71))
	big17 := big.NewInt(int64(17))
	big56 := big.NewInt(int64(56))

	a := FieldElement{big0, big223}
	b := FieldElement{big7, big223}
	x1 := FieldElement{big47, big223}
	y1 := FieldElement{big71, big223}
	x2 := FieldElement{big17, big223}
	y2 := FieldElement{big56, big223}
	p1 := Point{a, b, x1, y1}
	p2 := Point{a, b, x2, y2}
	t := new(FieldElement)
	//	fmt.Println(s.Add(s, x1).num)
	fmt.Println("when use new func:", t)
	fmt.Println("print fieldelement directly:", a)
	fmt.Println("p1.b.num", p1.b.num)
	fmt.Println("print p2 directly", p2)
	p3 := p1.AddPoint(p2)
	fmt.Println("p3.x.num:", p3.x.num)
	//	fmt.Println("p3.y.num:", p3.y.num)
}
