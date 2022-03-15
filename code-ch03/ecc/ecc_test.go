package ecc

import "testing"

func TestNe(t *testing.T) {
	a := NewFieldElement(2, 31)
	b := NewFieldElement(2, 31)
	c := NewFieldElement(15, 31)
	if !a.eq(b) {
		t.Fail()
	}
	if !a.ne(c) {
		t.Fail()
	}
	if a.ne(b) {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	a := NewFieldElement(2, 31)
	b := NewFieldElement(15, 31)
	if a.add(b).ne(NewFieldElement(17, 31)) {
		t.Fail()
	}
	a = NewFieldElement(17, 31)
	b = NewFieldElement(21, 31)
	if a.add(b).ne(NewFieldElement(7, 31)) {
		t.Fail()
	}
}

func TestMod(t *testing.T) {
	if mod(0, 1) != 0 {
		t.Fail()
	}
	if mod(0, 2) != 0 {
		t.Fail()
	}
	if mod(1, 2) != 1 {
		t.Fail()
	}
	if mod(2, 2) != 0 {
		t.Fail()
	}
	if mod(-1, 2) != 1 {
		t.Fail()
	}
}

func TestPowMod(t *testing.T) {
	if powMod(0, 0, 123) != 1 {
		t.Fail() // Slightly contraversial
	}
	if powMod(0, 1, 123) != 0 {
		t.Fail()
	}
	if powMod(1, 0, 123) != 1 {
		t.Fail()
	}
	if powMod(2, 2, 123) != 4 {
		t.Fail()
	}
	if powMod(4, 4, 5) != 1 {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	a := NewFieldElement(29, 31)
	b := NewFieldElement(4, 31)
	if a.sub(b).ne(NewFieldElement(25, 31)) {
		t.Fail()
	}
	a = NewFieldElement(15, 31)
	b = NewFieldElement(30, 31)
	if a.sub(b).ne(NewFieldElement(16, 31)) {
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	a := NewFieldElement(24, 31)
	b := NewFieldElement(19, 31)
	if a.mul(b).ne(NewFieldElement(22, 31)) {
		t.Fail()
	}
}

func TestPow(t *testing.T) {
	a := NewFieldElement(17, 31)
	if a.pow(3).ne(NewFieldElement(15, 31)) {
		t.Fail()
	}
	a = NewFieldElement(5, 31)
	b := NewFieldElement(18, 31)
	if a.pow(5).mul(b).ne(NewFieldElement(16, 31)) {
		t.Fail()
	}
}

func TestDiv(t *testing.T) {
	a := NewFieldElement(3, 31)
	b := NewFieldElement(24, 31)
	if a.div(b).ne(NewFieldElement(4, 31)) {
		t.Fail()
	}
	a = NewFieldElement(17, 31)
	if a.pow(-3).ne(NewFieldElement(29, 31)) {
		t.Fail()
	}
	a = NewFieldElement(4, 31)
	b = NewFieldElement(11, 31)
	if a.pow(-4).mul(b).ne(NewFieldElement(13, 31)) {
		t.Fail()
	}
}

/*
// These test are no longer valid now that NewPoint takes FieldElements
func TestPointNe(t *testing.T) {
	a := NewPoint(3, 7, 5, 7)
	b := NewPoint(18, 77, 5, 7)
	if a.ne(b) == false {
		t.Fail()
	}
	if a.ne(a) {
		t.Fail()
	}
}

func TestPointAdd0(t *testing.T) {
	a := NewInfPoint(5, 7)
	b := NewPoint(2, 5, 5, 7)
	c := NewPoint(2, -5, 5, 7)
	if a.add(b).ne(b) {
		t.Fail()
	}
	if b.add(a).ne(b) {
		t.Fail()
	}
	if b.add(c).ne(a) {
		t.Fail()
	}
}

func TestPointAdd1(t *testing.T) {
	a := NewPoint(3, 7, 5, 7)
	b := NewPoint(-1, -1, 5, 7)
	if a.add(b).ne(NewPoint(2, -5, 5, 7)) {
		t.Fail()
	}
}

func TestPointAdd2(t *testing.T) {
	a := NewPoint(-1, -1, 5, 7)
	if a.add(a).ne(NewPoint(18, 77, 5, 7)) {
		t.Fail()
	}
}
*/

func TestEccTestOnCurve(t *testing.T) {
	prime := 223
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	//https://golang.org/ref/spec#Composite_literals
	valid_points := [][]int{{192, 105}, {17, 56}, {1, 193}}
	//invalid_points := [][]int{{200, 119}, {42, 99}}
	for _, point := range valid_points {
		x := NewFieldElement(point[0], prime)
		y := NewFieldElement(point[1], prime)
		p := NewPoint(x, y, a, b)
		println(p.repr())
	}
}

// Answer Exercise 3
func TestEccAdd(t *testing.T) {
	prime := 223
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	p1 := NewPoint(NewFieldElement(170, prime), NewFieldElement(142, prime), a, b)
	p2 := NewPoint(NewFieldElement(60, prime), NewFieldElement(139, prime), a, b)
	sum := NewPoint(NewFieldElement(220, prime), NewFieldElement(181, prime), a, b)
	if p1.add(p2).ne(sum) {
		t.Fail()
	}
	p1 = NewPoint(NewFieldElement(47, prime), NewFieldElement(71, prime), a, b)
	p2 = NewPoint(NewFieldElement(17, prime), NewFieldElement(56, prime), a, b)
	sum = NewPoint(NewFieldElement(215, prime), NewFieldElement(68, prime), a, b)
	if p1.add(p2).ne(sum) {
		t.Fail()
	}
	p1 = NewPoint(NewFieldElement(143, prime), NewFieldElement(98, prime), a, b)
	p2 = NewPoint(NewFieldElement(76, prime), NewFieldElement(66, prime), a, b)
	sum = NewPoint(NewFieldElement(47, prime), NewFieldElement(71, prime), a, b)
	if p1.add(p2).ne(sum) {
		t.Fail()
	}
}
