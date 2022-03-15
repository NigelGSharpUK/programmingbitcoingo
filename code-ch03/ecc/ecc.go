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

func (fe *FieldElement) repr() string {
	return "FieldElement_" + strconv.Itoa(fe.prime) + "(" + strconv.Itoa(fe.num) + ")"
}

// Test for equality
func (fe *FieldElement) eq(other *FieldElement) bool {
	if fe == nil || other == nil {
		panic("Cannot compare nil pointers")
	}
	return fe.num == other.num && fe.prime == other.prime
}

// Test for inequality
func (fe *FieldElement) ne(other *FieldElement) bool {
	//panic("Not Implemented")

	// Answer Exercise 1
	if fe == nil || other == nil {
		panic("Cannot compare nil pointers")
	}
	return !fe.eq(other)
}

func (fe *FieldElement) add(other *FieldElement) *FieldElement {
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
func mod(a, b int) int {
	return (a%b + b) % b
}

func (fe *FieldElement) sub(other *FieldElement) *FieldElement {
	//panic("Not Implemented")

	// Answer Exercise 3
	if fe == nil || other == nil {
		panic("Cannot subtract nil pointers")
	}
	if fe.prime != other.prime {
		panic("Cannot subtract two numbers in different Fields")
	}
	num := mod((fe.num - other.num), fe.prime)
	return NewFieldElement(num, fe.prime)
}

func (fe *FieldElement) mul(other *FieldElement) *FieldElement {
	//panic("Not Implemented")

	// Answer Exercise 6
	if fe == nil || other == nil {
		panic("Cannot multiply nil pointers")
	}
	if fe.prime != other.prime {
		panic("Cannot multiply two numbers in different Fields")
	}
	num := mod((fe.num * other.num), fe.prime)
	return NewFieldElement(num, fe.prime)
}

// Need a pow function with modulus, like in Python
func powMod(base int, exp int, modulus int) int {
	if exp < 0 {
		panic("Negative exponent not supported here")
	}
	if exp == 0 {
		return 1
	} else if exp == 1 {
		return mod(base, modulus)
	} else {
		res := 1
		for i := 0; i < exp; i++ {
			res = mod(res*base, modulus)
		}
		return res
	}
}

func (fe *FieldElement) pow(exp int) *FieldElement {
	n := mod(exp, (fe.prime - 1))
	num := powMod(fe.num, n, fe.prime)
	return NewFieldElement(num, fe.prime)
}

func (fe *FieldElement) div(other *FieldElement) *FieldElement {
	//panic("Not Implemented")

	// Answer Exercise 9
	if fe == nil || other == nil {
		panic("Cannot divide nil pointers")
	}
	if fe.prime != other.prime {
		panic("Cannot divide two numbers in different Fields")
	}
	// Using Fermat's Little Theorem
	num := mod(fe.num*powMod(other.num, fe.prime-2, fe.prime), fe.prime)
	return NewFieldElement(num, fe.prime)
}

func (fe *FieldElement) rmul(coefficient int) *FieldElement {
	num := mod(fe.num*coefficient, fe.prime)
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
	if y.pow(2).ne(x.pow(3).add(a.mul(x)).add(b)) {
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

func (p *Point) eq(other *Point) bool {
	return p.x.eq(other.x) && p.y.eq(other.y) && p.a.eq(other.a) && p.b.eq(other.b)
}

func (p *Point) ne(other *Point) bool {
	//panic("Not Implemented")

	// Exercise 2 answer
	return !p.eq(other)
}

func (p *Point) repr() string {
	if p.isInf {
		return "Point(infinity)"
	}
	return "Point(" + p.x.repr() + "," + p.y.repr() + ")_" + p.a.repr() + "_" + p.b.repr()
}

func (p *Point) rmul(coefficient int) *Point {
	coef := coefficient
	current := p
	result := NewInfPoint(p.a, p.b) // Point at infinity acts as zero
	for coef != 0 {
		if coef&1 == 1 {
			result = result.add(current)
		}
		current = current.add(current)
		coef >>= 1
	}
	return result
}

func (p *Point) add(other *Point) *Point {
	if p.a.ne(other.a) || p.b.ne(other.b) {
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
	if p.eq(other) && p.y.num == 0 {
		return NewInfPoint(p.a, p.b)
	}

	// Case 1: self.x == other.x, self.y != other.y
	// Result is point at infinity
	// panic("Not implemented")

	// Answer Exercise 3
	if p.x.eq(other.x) && p.y.ne(other.y) {
		return NewInfPoint(p.a, p.b)
	}

	// Case 2: self.x ≠ other.x
	// Formula (x3,y3)==(x1,y1)+(x2,y2)
	// s=(y2-y1)/(x2-x1)
	// x3=s**2-x1-x2
	// y3=s*(x1-x3)-y1
	if p.x.ne(other.x) {
		// panic( "Not implemented")

		// Answer Exercise 5
		s := other.y.sub(p.y).div(other.x.sub(p.x))
		x3 := s.pow(2).sub(p.x).sub(other.x)
		y3 := s.mul(p.x.sub(x3)).sub(p.y)
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
	if p.x.eq(other.x) && p.y.eq(other.y) {
		s := p.x.pow(2).rmul(3).add(p.a).div(p.y.rmul(2))
		x3 := s.pow(2).sub(p.x.rmul(2))
		y3 := s.mul(p.x.sub(x3)).sub(p.y)
		return NewPoint(x3, y3, p.a, p.b)
	}

	// Final case, tangent is vertical
	if p.x == other.x && p.y.num == 0 {
		return NewInfPoint(p.a, p.b)
	}

	panic("Fell through, missing case?")
}
