package ecc

import "testing"

func TestNeq(t *testing.T) {
	a, _ := NewFieldElement(2, 31)
	b, _ := NewFieldElement(2, 31)
	c, _ := NewFieldElement(15, 31)
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
	a, _ := NewFieldElement(2, 31)
	b, _ := NewFieldElement(15, 31)
	c, _ := NewFieldElement(17, 31)
	if a.add(b).ne(c) {
		t.Fail()
	}
	a, _ = NewFieldElement(17, 31)
	b, _ = NewFieldElement(21, 31)
	c, _ = NewFieldElement(7, 31)
	if a.add(b).ne(c) {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	a, _ := NewFieldElement(29, 31)
	b, _ := NewFieldElement(4, 31)
	c, _ := NewFieldElement(25, 31)
	if a.sub(b).ne(c) {
		t.Fail()
	}
	a, _ = NewFieldElement(15, 31)
	b, _ = NewFieldElement(30, 31)
	c, _ = NewFieldElement(16, 31)
	if a.sub(b).ne(c) {
		t.Fail()
	}
}
