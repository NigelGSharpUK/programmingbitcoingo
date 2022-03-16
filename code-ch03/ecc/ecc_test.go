package ecc

import (
	"math/big"
	"testing"
)

func TestEq(t *testing.T) {
	a := NewFieldElement(2, 31)
	b := NewFieldElement(2, 31)
	c := NewFieldElement(15, 31)
	if !a.Eq(b) {
		t.Fail()
	}
	if a.Eq(c) {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	a := NewFieldElement(2, 31)
	b := NewFieldElement(15, 31)
	if !a.Add(b).Eq(NewFieldElement(17, 31)) {
		t.Fail()
	}
	a = NewFieldElement(17, 31)
	b = NewFieldElement(21, 31)
	if !a.Add(b).Eq(NewFieldElement(7, 31)) {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	a := NewFieldElement(29, 31)
	b := NewFieldElement(4, 31)
	if !a.Sub(b).Eq(NewFieldElement(25, 31)) {
		t.Fail()
	}
	a = NewFieldElement(15, 31)
	b = NewFieldElement(30, 31)
	if !a.Sub(b).Eq(NewFieldElement(16, 31)) {
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	a := NewFieldElement(24, 31)
	b := NewFieldElement(19, 31)
	if !a.Mul(b).Eq(NewFieldElement(22, 31)) {
		t.Fail()
	}
}

func TestPow(t *testing.T) {
	a := NewFieldElement(17, 31)
	if !a.Pow(big.NewInt(3)).Eq(NewFieldElement(15, 31)) {
		t.Fail()
	}
	a = NewFieldElement(5, 31)
	b := NewFieldElement(18, 31)
	if !a.Pow(big.NewInt(5)).Mul(b).Eq(NewFieldElement(16, 31)) {
		t.Fail()
	}
}

func TestDiv(t *testing.T) {
	a := NewFieldElement(3, 31)
	b := NewFieldElement(24, 31)
	if !a.Div(b).Eq(NewFieldElement(4, 31)) {
		t.Fail()
	}
	a = NewFieldElement(17, 31)
	if !a.Pow(big.NewInt(-3)).Eq(NewFieldElement(29, 31)) {
		t.Fail()
	}
	a = NewFieldElement(4, 31)
	b = NewFieldElement(11, 31)
	if !a.Pow(big.NewInt(-4)).Mul(b).Eq(NewFieldElement(13, 31)) {
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
	prime := int64(223)
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	//https://golang.org/ref/spec#Composite_literals
	valid_points := [][]int64{{192, 105}, {17, 56}, {1, 193}}
	//invalid_points := [][]int{{200, 119}, {42, 99}}
	for _, point := range valid_points {
		x := NewFieldElement(point[0], prime)
		y := NewFieldElement(point[1], prime)
		p := NewPoint(x, y, a, b)
		println(p.Repr()) // [  ] Get rid of this print and do the test from the book properly
	}
}

// Answer Exercise 3
func TestEccAdd(t *testing.T) {
	prime := int64(223)
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	p1 := NewPoint(NewFieldElement(170, prime), NewFieldElement(142, prime), a, b)
	p2 := NewPoint(NewFieldElement(60, prime), NewFieldElement(139, prime), a, b)
	sum := NewPoint(NewFieldElement(220, prime), NewFieldElement(181, prime), a, b)
	if !p1.Add(p2).Eq(sum) {
		t.Fail()
	}
	p1 = NewPoint(NewFieldElement(47, prime), NewFieldElement(71, prime), a, b)
	p2 = NewPoint(NewFieldElement(17, prime), NewFieldElement(56, prime), a, b)
	sum = NewPoint(NewFieldElement(215, prime), NewFieldElement(68, prime), a, b)
	if !p1.Add(p2).Eq(sum) {
		t.Fail()
	}
	p1 = NewPoint(NewFieldElement(143, prime), NewFieldElement(98, prime), a, b)
	p2 = NewPoint(NewFieldElement(76, prime), NewFieldElement(66, prime), a, b)
	sum = NewPoint(NewFieldElement(47, prime), NewFieldElement(71, prime), a, b)
	if !p1.Add(p2).Eq(sum) {
		t.Fail()
	}
}

// We must be able to compare with inf (we found that we couldn't at first!)
func TestInfCompare(t *testing.T) {
	prime := int64(223)
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	infx := NewInfPoint(a, b)
	infy := NewInfPoint(a, b)
	if !infx.Eq(infy) {
		t.Fail()
	}
	x := NewFieldElement(15, prime)
	y := NewFieldElement(86, prime)
	z := NewPoint(x, y, a, b)
	if infx.Eq(z) {
		t.Fail()
	}
	if z.Eq(infy) {
		t.Fail()
	}
}

func TestBitcoinParams(t *testing.T) {
	// Verify Gx,Gy are on the curve y^2 = x^3 +ax +b (with a,b = 0,7)
	bitcoinParams := secp256k1_Params()
	Gx := &bitcoinParams.Gx
	Gy := &bitcoinParams.Gy
	a := &bitcoinParams.a
	b := &bitcoinParams.b
	p := &bitcoinParams.p
	var ysquared big.Int
	ysquared.Exp(Gy, big.NewInt(2), p)
	var xcubed big.Int
	xcubed.Exp(Gx, big.NewInt(3), p)
	var ax big.Int
	ax.Mul(a, Gx)
	lhsModP := ysquared // Already Mod'd by the parameter p in Exp() above
	var xcubedPlusAx big.Int
	xcubedPlusAx.Add(&xcubed, &ax)
	var rhs big.Int
	rhs.Add(&xcubedPlusAx, b)
	var rhsModP big.Int
	rhsModP.Mod(&rhs, p)
	//fmt.Printf("0x%064X\n", &lhsModP)
	//fmt.Printf("0x%064X\n", &rhsModP)
	if lhsModP.Cmp(&rhsModP) != 0 {
		t.Fail()
	}

	// Verify the order of the group generated by point Gx,Gy is n
	G := G()              // Bitcoin's generator point
	n := &bitcoinParams.n // Bitcoin's order of the group generated by G
	inf := G.Rmul(n)      // Should be the zero point (at infinity)
	if !inf.Eq(NewS256InfPoint()) {
		t.Fail()
	}
}
