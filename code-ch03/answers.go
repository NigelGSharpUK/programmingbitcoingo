package main

import (
	"math/big"
	. "programmingbitcoingo/code-ch03/ecc"
)

func main() {
	exercise1()
	exercise2()
	println("Exercise 3 is in ecc_test.go")
	exercise4()
	exercise5()
}

func on_curve(x, y, a, b *FieldElement) bool {
	//return y^2 == x^3 +ax +b
	lhs := y.Pow(big.NewInt(2))
	xcubed := x.Pow(big.NewInt(3))
	var ax FieldElement
	ax.Mul(a, x)
	var xcubedPlusAx FieldElement
	xcubedPlusAx.Add(xcubed, &ax)
	var rhs FieldElement
	rhs.Add(&xcubedPlusAx, b)
	return lhs.Eq(&rhs)
}

func exercise1() {
	println("Exercise 1")
	prime := int64(223)
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	println(on_curve(NewFieldElement(192, prime), NewFieldElement(105, prime), a, b))
	println(on_curve(NewFieldElement(17, prime), NewFieldElement(56, prime), a, b))
	println(on_curve(NewFieldElement(200, prime), NewFieldElement(119, prime), a, b))
	println(on_curve(NewFieldElement(1, prime), NewFieldElement(193, prime), a, b))
	println(on_curve(NewFieldElement(42, prime), NewFieldElement(99, prime), a, b))
}

func exercise2() {
	println("Exercise 2")
	prime := int64(223)
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	x := NewPoint(NewFieldElement(170, prime), NewFieldElement(142, prime), a, b)
	y := NewPoint(NewFieldElement(60, prime), NewFieldElement(139, prime), a, b)
	println(x.Add(y).Repr())
	x = NewPoint(NewFieldElement(47, prime), NewFieldElement(71, prime), a, b)
	y = NewPoint(NewFieldElement(17, prime), NewFieldElement(56, prime), a, b)
	println(x.Add(y).Repr())
	x = NewPoint(NewFieldElement(143, prime), NewFieldElement(98, prime), a, b)
	y = NewPoint(NewFieldElement(76, prime), NewFieldElement(66, prime), a, b)
	println(x.Add(y).Repr())
}

func exercise4() {
	println("Exercise 4")
	prime := int64(223)
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	x := NewPoint(NewFieldElement(192, prime), NewFieldElement(105, prime), a, b)
	println(x.Add(x).Repr())
	x = NewPoint(NewFieldElement(143, prime), NewFieldElement(98, prime), a, b)
	println(x.Add(x).Repr())
	x = NewPoint(NewFieldElement(47, prime), NewFieldElement(71, prime), a, b)
	x = x.Add(x)
	println(x.Repr()) // *2
	x = x.Add(x)
	println(x.Repr()) // *4
	x = x.Add(x)
	println(x.Repr()) // *8
	origx := NewPoint(NewFieldElement(47, prime), NewFieldElement(71, prime), a, b)
	x = NewInfPoint(a, b)
	for i := 0; i < 21; i++ {
		x = x.Add(origx)
	}
	println(x.Repr()) // *21
}

func exercise5() {
	println("Exercise 5")
	prime := int64(223)
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	origx := NewPoint(NewFieldElement(15, prime), NewFieldElement(86, prime), a, b)
	inf := NewInfPoint(a, b)
	x := NewInfPoint(a, b)
	var order int
	for i := 1; true; i++ {
		x = x.Add(origx)
		if x.Eq(inf) {
			order = i
			break
		}
	}
	println("Order of group generated:", order)
}
