package ecc

import "testing"

func TestNe(t *testing.T) {
	a := NewFieldElement(2, 31)
	b := NewFieldElement(2, 31)
	c := NewFieldElement(15, 31)
	if !a.Eq(b) {
		t.Fail()
	}
	if !a.Ne(c) {
		t.Fail()
	}
	if a.Ne(b) {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	a := NewFieldElement(2, 31)
	b := NewFieldElement(15, 31)
	if a.Add(b).Ne(NewFieldElement(17, 31)) {
		t.Fail()
	}
	a = NewFieldElement(17, 31)
	b = NewFieldElement(21, 31)
	if a.Add(b).Ne(NewFieldElement(7, 31)) {
		t.Fail()
	}
}

func TestMod(t *testing.T) {
	if Mod(0, 1) != 0 {
		t.Fail()
	}
	if Mod(0, 2) != 0 {
		t.Fail()
	}
	if Mod(1, 2) != 1 {
		t.Fail()
	}
	if Mod(2, 2) != 0 {
		t.Fail()
	}
	if Mod(-1, 2) != 1 {
		t.Fail()
	}
}

func TestPowMod(t *testing.T) {
	if PowMod(0, 0, 123) != 1 {
		t.Fail() // Slightly contraversial
	}
	if PowMod(0, 1, 123) != 0 {
		t.Fail()
	}
	if PowMod(1, 0, 123) != 1 {
		t.Fail()
	}
	if PowMod(2, 2, 123) != 4 {
		t.Fail()
	}
	if PowMod(4, 4, 5) != 1 {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	a := NewFieldElement(29, 31)
	b := NewFieldElement(4, 31)
	if a.Sub(b).Ne(NewFieldElement(25, 31)) {
		t.Fail()
	}
	a = NewFieldElement(15, 31)
	b = NewFieldElement(30, 31)
	if a.Sub(b).Ne(NewFieldElement(16, 31)) {
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	a := NewFieldElement(24, 31)
	b := NewFieldElement(19, 31)
	if a.Mul(b).Ne(NewFieldElement(22, 31)) {
		t.Fail()
	}
}

func TestPow(t *testing.T) {
	a := NewFieldElement(17, 31)
	if a.Pow(3).Ne(NewFieldElement(15, 31)) {
		t.Fail()
	}
	a = NewFieldElement(5, 31)
	b := NewFieldElement(18, 31)
	if a.Pow(5).Mul(b).Ne(NewFieldElement(16, 31)) {
		t.Fail()
	}
}

func TestDiv(t *testing.T) {
	a := NewFieldElement(3, 31)
	b := NewFieldElement(24, 31)
	if a.Div(b).Ne(NewFieldElement(4, 31)) {
		t.Fail()
	}
	a = NewFieldElement(17, 31)
	if a.Pow(-3).Ne(NewFieldElement(29, 31)) {
		t.Fail()
	}
	a = NewFieldElement(4, 31)
	b = NewFieldElement(11, 31)
	if a.Pow(-4).Mul(b).Ne(NewFieldElement(13, 31)) {
		t.Fail()
	}
}

func TestPointNe(t *testing.T) {
	a := NewPoint(3, 7, 5, 7)
	b := NewPoint(18, 77, 5, 7)
	if a.Ne(b) == false {
		t.Fail()
	}
	if a.Ne(a) {
		t.Fail()
	}
}

func TestPointAdd0(t *testing.T) {
	a := NewInfPoint(5, 7)
	b := NewPoint(2, 5, 5, 7)
	c := NewPoint(2, -5, 5, 7)
	if a.Add(b).Ne(b) {
		t.Fail()
	}
	if b.Add(a).Ne(b) {
		t.Fail()
	}
	if b.Add(c).Ne(a) {
		t.Fail()
	}
}

func TestPointAdd1(t *testing.T) {
	a := NewPoint(3, 7, 5, 7)
	b := NewPoint(-1, -1, 5, 7)
	if a.Add(b).Ne(NewPoint(2, -5, 5, 7)) {
		t.Fail()
	}
}

func TestPointAdd2(t *testing.T) {
	a := NewPoint(-1, -1, 5, 7)
	if a.Add(a).Ne(NewPoint(18, 77, 5, 7)) {
		t.Fail()
	}
}
