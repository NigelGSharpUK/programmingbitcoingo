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
