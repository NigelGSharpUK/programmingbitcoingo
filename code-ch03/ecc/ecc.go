package ecc

import (
	"fmt"
	"math/big"
)

type CurveParams struct {
	a  *big.Int
	b  *big.Int
	p  *big.Int
	Gx *big.Int
	Gy *big.Int
	n  *big.Int
}

func secp256k1_Params() *CurveParams {
	params := new(CurveParams)
	params.a = big.NewInt(0)
	params.b = big.NewInt(7)
	// p = 2^256 - 2^32 - 977
	var p2_256 big.Int
	p2_256.Exp(big.NewInt(2), big.NewInt(256), nil)
	var p2_32 big.Int
	p2_32.Exp(big.NewInt(2), big.NewInt(32), nil)
	var diffPowers2 big.Int
	diffPowers2.Sub(&p2_256, &p2_32)
	params.p = big.NewInt(0)
	params.p.Sub(&diffPowers2, big.NewInt(977))
	// Gx = a big hex number
	params.Gx = big.NewInt(0)
	params.Gx.SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
	// Gy = a big hex number
	params.Gy = big.NewInt(0)
	params.Gy.SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)
	// n = a big hex number
	params.n = big.NewInt(0)
	params.n.SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	return params
}

type FieldElement struct {
	num   *big.Int
	prime *big.Int
}

func NewFieldElement(num *big.Int, prime *big.Int) *FieldElement {
	//if num >= prime || num < 0 {
	if num.Cmp(prime) >= 0 || num.Cmp(big.NewInt(0)) == -1 {
		panic("num must be between 0 and prime-1 inclusive")
	}
	fe := new(FieldElement)
	fe.num = num
	fe.prime = prime
	return fe
}

// Convenience fn without needing big.NewInt() all over the place
func NewFieldElement_(num int, prime int) *FieldElement {
	bigNum := big.NewInt(int64(num))
	bigPrime := big.NewInt(int64(prime))
	return NewFieldElement(bigNum, bigPrime)
}

func (fe *FieldElement) Repr() string {
	primeOrder := ""
	if fe.prime.Cmp(secp256k1_Params().p) == 0 {
		primeOrder = "secp256k1"
	} else {
		primeOrder = fmt.Sprintf("%d", fe.prime)
	}
	val := ""
	if fe.num.Cmp(big.NewInt(65536)) == -1 {
		val = fmt.Sprintf("%d", fe.num) // Small numbers as decimal
	} else {
		val = fmt.Sprintf("%064X", fe.num) // Big numbers as 256 bit hex
	}

	return "FieldElement_" + primeOrder + "(" + val + ")"
}

// Test for equality
func (fe *FieldElement) Eq(other *FieldElement) bool {
	if fe == nil || other == nil {
		panic("Cannot compare nil pointers")
	}
	//return fe.num == other.num && fe.prime == other.prime
	return fe.num.Cmp(other.num) == 0 && fe.prime.Cmp(other.prime) == 0
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
	//if fe.prime != other.prime {
	if fe.prime.Cmp(other.prime) != 0 {
		panic("Cannot add two numbers in different Fields")
	}
	var num big.Int
	num.Add(fe.num, other.num)
	num.Mod(&num, fe.prime) // [  ] I think this is OK - Euclidean Modulus unlike Go
	return NewFieldElement(&num, fe.prime)
}

func (fe *FieldElement) Sub(other *FieldElement) *FieldElement {
	//panic("Not Implemented")

	// Answer Exercise 3
	if fe == nil || other == nil {
		panic("Cannot subtract nil pointers")
	}
	//if fe.prime != other.prime {
	if fe.prime.Cmp(other.prime) != 0 {
		panic("Cannot subtract two numbers in different Fields")
	}
	//num := Mod((fe.num - other.num), fe.prime)
	var num big.Int
	num.Sub(fe.num, other.num)
	num.Mod(&num, fe.prime)
	return NewFieldElement(&num, fe.prime)
}

func (fe *FieldElement) Mul(other *FieldElement) *FieldElement {
	//panic("Not Implemented")

	// Answer Exercise 6
	if fe == nil || other == nil {
		panic("Cannot multiply nil pointers")
	}
	//if fe.prime != other.prime {
	if fe.prime.Cmp(other.prime) != 0 {
		panic("Cannot multiply two numbers in different Fields")
	}
	//num := Mod((fe.num * other.num), fe.prime)
	var num big.Int
	num.Mul(fe.num, other.num)
	num.Mod(&num, fe.prime)
	return NewFieldElement(&num, fe.prime)
}

func (fe *FieldElement) Pow(exp *big.Int) *FieldElement {
	//n := Mod(exp, (fe.prime - 1))		[  ] Really? -1?
	var primeMinusOne big.Int
	primeMinusOne.Sub(fe.prime, big.NewInt(1))
	var n big.Int
	n.Mod(exp, &primeMinusOne)
	//num := PowMod(fe.num, n, fe.prime)
	var num big.Int
	num.Exp(fe.num, &n, fe.prime)
	return NewFieldElement(&num, fe.prime)
}

func (fe *FieldElement) Div(other *FieldElement) *FieldElement {
	//panic("Not Implemented")

	// Answer Exercise 9
	if fe == nil || other == nil {
		panic("Cannot divide nil pointers")
	}
	//if fe.prime != other.prime {
	if fe.prime.Cmp(other.prime) != 0 {
		panic("Cannot divide two numbers in different Fields")
	}
	// Using Fermat's Little Theorem
	//num := Mod(fe.num*PowMod(other.num, fe.prime-2, fe.prime), fe.prime)
	var pMinusTwo big.Int
	pMinusTwo.Sub(fe.prime, big.NewInt(2))
	var numToPMinusTwo big.Int
	numToPMinusTwo.Exp(other.num, &pMinusTwo, fe.prime)
	var product big.Int
	product.Mul(fe.num, &numToPMinusTwo)
	var num big.Int
	num.Mod(&product, fe.prime)
	return NewFieldElement(&num, fe.prime)
}

func (fe *FieldElement) Rmul(coefficient *big.Int) *FieldElement {
	//num := Mod(fe.num*coefficient, fe.prime)
	var product big.Int
	product.Mul(fe.num, coefficient)
	var num big.Int
	num.Mod(&product, fe.prime)
	return NewFieldElement(&num, fe.prime)
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
	//if y.Pow(2).Ne(x.Pow(3).Add(a.Mul(x)).Add(b)) {
	if y.Pow(big.NewInt(2)).Ne(x.Pow(big.NewInt(3)).Add(a.Mul(x)).Add(b)) {
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
	// First check for both infinities
	if p.isInf && other.isInf {
		return p.a.Eq(other.a) && p.b.Eq(other.b)
	}
	// Then check for one or other infinities
	if p.isInf || other.isInf {
		// We know they're not both infinity
		return false
	}
	// Now we're safe to compare FiniteFields
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
	// Print over two lines
	return "Point(" + p.x.Repr() + ",\n      " + p.y.Repr() + ")_" + p.a.Repr() + "_" + p.b.Repr()
}

func (p *Point) Rmul(coefficient *big.Int) *Point {
	coef := coefficient // [  ] OK? Or are we just copying pointers?!
	current := p
	result := NewInfPoint(p.a, p.b) // Point at infinity acts as zero
	//for coef != 0 {
	for coef.Cmp(big.NewInt(0)) != 0 {
		//if coef&1 == 1 {
		var lsb big.Int
		lsb.And(coef, big.NewInt(1))
		if lsb.Cmp(big.NewInt(1)) == 0 {
			result = result.Add(current)
		}
		current = current.Add(current)
		//coef >>= 1
		coef.Rsh(coef, 1)
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
	//if p.Eq(other) && p.y.num == 0 {
	if p.Eq(other) && p.y.num.Cmp(big.NewInt(0)) == 0 {
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
		x3 := s.Pow(big.NewInt(2)).Sub(p.x).Sub(other.x)
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
		s := p.x.Pow(big.NewInt(2)).Rmul(big.NewInt(3)).Add(p.a).Div(p.y.Rmul(big.NewInt(2)))
		x3 := s.Pow(big.NewInt(2)).Sub(p.x.Rmul(big.NewInt(2)))
		y3 := s.Mul(p.x.Sub(x3)).Sub(p.y)
		return NewPoint(x3, y3, p.a, p.b)
	}

	// Final case, tangent is vertical
	if p.x == other.x && p.y.num.Cmp(big.NewInt(0)) == 0 {
		return NewInfPoint(p.a, p.b)
	}

	panic("Fell through, missing case?")
}
