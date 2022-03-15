package ecc

import (
	"strconv"
)

type FieldElement struct {
	num   int // [  ] Someday soon we'll need numbers bigger than int
	prime int
}

func NewFieldElement(num int, prime int) *FieldElement {
	if num >= prime || num < 0 {
		panic("num must be between 0 and prime-1 inclusive")
	}
	fe := new(FieldElement)
	fe.num = num
	fe.prime = prime
	return fe
}

func (fe *FieldElement) Repr() string {
	return "FieldElement_" + strconv.Itoa(fe.prime) + "(" + strconv.Itoa(fe.num) + ")"
}

// Test for equality
func (fe *FieldElement) Eq(other *FieldElement) bool {
	if fe == nil || other == nil {
		panic("Cannot compare nil pointers")
	}
	return fe.num == other.num && fe.prime == other.prime
}

// Test for inequality
func (fe *FieldElement) Ne(other *FieldElement) bool {
	//panic("Not Implemented")

	// Answer Exercise 1
	if fe == nil || other == nil {
		panic("Cannot compare nil pointers")
	}
	return !fe.Eq(other)
}

func (fe *FieldElement) Add(other *FieldElement) *FieldElement {
	if fe == nil || other == nil {
		panic("Cannot add nil pointers")
	}
	if fe.prime != other.prime {
		panic("Cannot add two numbers in different Fields")
	}
	num := (fe.num + other.num) % fe.prime // [  ] Warning % only works like Python for +ve num
	return NewFieldElement(num, fe.prime)
}

// Go's % operator is DIFFERENT to Python's % operator
func Mod(a, b int) int {
	return (a%b + b) % b
}

func (fe *FieldElement) Sub(other *FieldElement) *FieldElement {
	//panic("Not Implemented")

	// Answer Exercise 3
	if fe == nil || other == nil {
		panic("Cannot subtract nil pointers")
	}
	if fe.prime != other.prime {
		panic("Cannot subtract two numbers in different Fields")
	}
	num := Mod((fe.num - other.num), fe.prime)
	return NewFieldElement(num, fe.prime)
}

func (fe *FieldElement) Mul(other *FieldElement) *FieldElement {
	//panic("Not Implemented")

	// Answer Exercise 6
	if fe == nil || other == nil {
		panic("Cannot multiply nil pointers")
	}
	if fe.prime != other.prime {
		panic("Cannot multiply two numbers in different Fields")
	}
	num := Mod((fe.num * other.num), fe.prime)
	return NewFieldElement(num, fe.prime)
}

// Need a pow function with modulus, like in Python
func PowMod(base int, exp int, modulus int) int {
	if exp < 0 {
		panic("Negative exponent not supported here")
	}
	if exp == 0 {
		return 1
	} else if exp == 1 {
		return Mod(base, modulus)
	} else {
		res := 1
		for i := 0; i < exp; i++ {
			res = Mod(res*base, modulus)
		}
		return res
	}
}

func (fe *FieldElement) Pow(exp int) *FieldElement {
	n := Mod(exp, (fe.prime - 1))
	num := PowMod(fe.num, n, fe.prime)
	return NewFieldElement(num, fe.prime)
}

func (fe *FieldElement) Div(other *FieldElement) *FieldElement {
	//panic("Not Implemented")

	// Answer Exercise 9
	if fe == nil || other == nil {
		panic("Cannot divide nil pointers")
	}
	if fe.prime != other.prime {
		panic("Cannot divide two numbers in different Fields")
	}
	// Using Fermat's Little Theorem
	num := Mod(fe.num*PowMod(other.num, fe.prime-2, fe.prime), fe.prime)
	return NewFieldElement(num, fe.prime)
}

func (fe *FieldElement) Rmul(coefficient int) *FieldElement {
	num := Mod(fe.num*coefficient, fe.prime)
	return NewFieldElement(num, fe.prime)
}

type Point struct {
	isInf bool
	x     *FieldElement // ignore if isInf
	y     *FieldElement // ignore if isInf
	a     *FieldElement
	b     *FieldElement
}

func NewPoint(x, y, a, b *FieldElement) *Point {
	res := new(Point)
	res.isInf = false
	res.x = x
	res.y = y
	res.a = a
	res.b = b
	if y.Pow(2).Ne(x.Pow(3).Add(a.Mul(x)).Add(b)) {
		panic("Point is not on the curve")
	}
	return res
}

func NewInfPoint(a, b *FieldElement) *Point {
	res := new(Point)
	res.isInf = true
	res.a = a
	res.b = b
	return res
}

func (p *Point) Eq(other *Point) bool {
	return p.x.Eq(other.x) && p.y.Eq(other.y) && p.a.Eq(other.a) && p.b.Eq(other.b)
}

func (p *Point) Ne(other *Point) bool {
	//panic("Not Implemented")

	// Exercise 2 answer
	return !p.Eq(other)
}

func (p *Point) Repr() string {
	if p.isInf {
		return "Point(infinity)"
	}
	return "Point(" + p.x.Repr() + "," + p.y.Repr() + ")_" + p.a.Repr() + "_" + p.b.Repr()
}

func (p *Point) Rmul(coefficient int) *Point {
	coef := coefficient
	current := p
	result := NewInfPoint(p.a, p.b) // Point at infinity acts as zero
	for coef != 0 {
		if coef&1 == 1 {
			result = result.Add(current)
		}
		current = current.Add(current)
		coef >>= 1
	}
	return result
}

func (p *Point) Add(other *Point) *Point {
	if p.a.Ne(other.a) || p.b.Ne(other.b) {
		panic("Can't add points that are not on same curve")
	}
	if p.isInf {
		// If p is point at infinity, it is the identity under addition
		return other // [  ] Do we need to make a copy?
	}
	if other.isInf {
		// If other is point at infinity, it is the identity under addition
		return p // [  ] Do we need to make a copy?
	}

	// Handle p==other and y==0 (vertical tangent)
	if p.Eq(other) && p.y.num == 0 {
		return NewInfPoint(p.a, p.b)
	}

	// Case 1: self.x == other.x, self.y != other.y
	// Result is point at infinity
	// panic("Not implemented")

	// Answer Exercise 3
	if p.x.Eq(other.x) && p.y.Ne(other.y) {
		return NewInfPoint(p.a, p.b)
	}

	// Case 2: self.x ≠ other.x
	// Formula (x3,y3)==(x1,y1)+(x2,y2)
	// s=(y2-y1)/(x2-x1)
	// x3=s**2-x1-x2
	// y3=s*(x1-x3)-y1
	if p.x.Ne(other.x) {
		// panic( "Not implemented")

		// Answer Exercise 5
		s := other.y.Sub(p.y).Div(other.x.Sub(p.x))
		x3 := s.Pow(2).Sub(p.x).Sub(other.x)
		y3 := s.Mul(p.x.Sub(x3)).Sub(p.y)
		return NewPoint(x3, y3, p.a, p.b)
	}

	// Case 3: self == other
	// Formula (x3,y3)=(x1,y1)+(x1,y1)
	// s=(3*x1**2+a)/(2*y1)
	// x3=s**2-2*x1
	// y3=s*(x1-x3)-y1
	// panic("Not implemented")

	// Answer Exercise 7
	// Handle p,other being same point, so use tangent
	if p.x.Eq(other.x) && p.y.Eq(other.y) {
		s := p.x.Pow(2).Rmul(3).Add(p.a).Div(p.y.Rmul(2))
		x3 := s.Pow(2).Sub(p.x.Rmul(2))
		y3 := s.Mul(p.x.Sub(x3)).Sub(p.y)
		return NewPoint(x3, y3, p.a, p.b)
	}

	// Final case, tangent is vertical
	if p.x == other.x && p.y.num == 0 {
		return NewInfPoint(p.a, p.b)
	}

	panic("Fell through, missing case?")
}
