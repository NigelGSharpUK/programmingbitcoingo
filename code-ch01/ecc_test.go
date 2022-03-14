package ecc

import "testing"

func TestNeq(t *testing.T) {
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
