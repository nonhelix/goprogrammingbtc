package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

//Page 4.
type FieldElement struct {
	num   int
	prime int
}

// TODO: check the valdiation of FieldElement, check the big int, Only the exp had big int compatiable now.

func add(x, y FieldElement) (FieldElement, error) {
	if x.prime == y.prime {
		// modulo
		a := (x.num + y.num) % x.prime
		z := FieldElement{a, x.prime}
		return z, nil
	} else {
		//TODO: check how to return illegal fieldelement, not {-1,-1}
		return FieldElement{-1, -1}, errors.New("can't work with two numbers in different Fields")
	}
}

func sub(x, y FieldElement) (FieldElement, error) {
	if x.prime == y.prime {
		// modulo
		a := (x.num - y.num) % x.prime
		z := FieldElement{a, x.prime}
		return z, nil
	} else {
		//TODO: check how to return illegal fieldelement, not {-1,-1}
		return FieldElement{-1, -1}, errors.New("can't work with two numbers in different Fields")
	}
}

func mul(x, y FieldElement) (FieldElement, error) {
	if x.prime == y.prime {
		// modulo
		a := (x.num * y.num) % x.prime
		z := FieldElement{a, x.prime}
		return z, nil
	} else {
		//TODO: check how to return illegal fieldelement, not {-1,-1}
		return FieldElement{-1, -1}, errors.New("can't work with two numbers in different Fields")
	}
}

// power had negative issues and big int issues TODO TODO TODO.

func pow(x FieldElement, exponent int) (FieldElement, error) {
	// modulo, in case exponenet is negative.
	e := exponent % (x.prime - 1)
	big_a := big.NewInt(int64(x.num))
	big_e := big.NewInt(int64(e))
	big_p := big.NewInt(int64(x.prime))
	big_a.Exp(big_a, big_e, big_p)
	small_a := int(big_a.Int64())
	z := FieldElement{small_a, x.prime}
	return z, nil
}

func trueDiv(x, y FieldElement) (FieldElement, error) {
	if x.prime == y.prime {
		// modulo
		y_fermat := math.Pow(float64(y.num), float64(y.prime-2))
		a := (x.num * int(y_fermat)) % x.prime
		z := FieldElement{a, x.prime}
		return z, nil
	} else {
		//TODO: check how to return illegal fieldelement, not {-1,-1}
		return FieldElement{-1, -1}, errors.New("can't work with two numbers in different Fields")
	}
}

//point

type Point struct {
	a int
	b int
	x int
	y int
}

// y2 = x3 + 5x + 7
func (p Point) onCurve() bool {
	return p.y*p.y == int(math.Pow(float64(p.x), 3))+5*p.x+7
}

func (p1 Point) addPoint(p2 Point) (p3 Point) {
	//TODO check onCurve if both are Inf are also OK.(pg.34) and check sameCurve

	// same x and not same y case
	if p1.x == p2.x && p1.y != p2.y {
		return Point{p1.a, p1.b, int(math.Inf(1)), int(math.Inf(1))}
	}
	// Inf case TODO Inf is not big enough. Maybe all of them shouble be bytes?
	if p1.x == int(math.Inf(1)) {
		return p2
	} else if p2.x == int(math.Inf(1)) {
		return p1
	}
	//x not equal
	if p1.x != p2.x {
		s := (p2.y - p1.y) / (p2.x - p1.x)
		x := s*s - p1.x - p2.x
		y := s*(p1.x-x) - p1.y
		return Point{p1.a, p1.b, x, y}
	}

	// P1 == P2 and y = 0
	if p1 == p2 && p1.y == 0 {
		return Point{p1.a, p1.b, int(math.Inf(1)), int(math.Inf(1))}
	}
	// P1 == P2
	if p1 == p2 {
		s := (3*(p1.x*p1.x) + p1.a) / (2 * p1.y)
		x := s*s - 2*p1.x
		y := s*(p1.x-x) - p1.y
		return Point{p1.a, p1.b, x, y}
	}
	return
}

func main() {
	//a := FieldElement{17, 31}
	//b := FieldElement{7, 19}
	//c, _ := pow(a, -3)
	//fmt.Println(c)

	a := Point{5, 7, -1, 0}

	b := Point{5, 7, -1, 0}
	//	c := Point{5, 7, 18, 77}
	//	d := Point{int(math.Inf(1)), int(math.Inf(1)), 5, 7}

	fmt.Println(a.addPoint(b))

}
