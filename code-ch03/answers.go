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
	var lhs FieldElement
	lhs.Exp(y, big.NewInt(2))
	var xcubed FieldElement
	xcubed.Exp(x, big.NewInt(3))
	var ax FieldElement
	ax.Mul(a, x)
	var xcubedPlusAx FieldElement
	xcubedPlusAx.Add(&xcubed, &ax)
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
	var sum Point
	sum.Add(x, y)
	println(sum.Repr())
	x = NewPoint(NewFieldElement(47, prime), NewFieldElement(71, prime), a, b)
	y = NewPoint(NewFieldElement(17, prime), NewFieldElement(56, prime), a, b)
	sum.Add(x, y)
	println(sum.Repr())
	x = NewPoint(NewFieldElement(143, prime), NewFieldElement(98, prime), a, b)
	y = NewPoint(NewFieldElement(76, prime), NewFieldElement(66, prime), a, b)
	sum.Add(x, y)
	println(sum.Repr())
}

func exercise4() {
	println("Exercise 4")
	prime := int64(223)
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	x := NewPoint(NewFieldElement(192, prime), NewFieldElement(105, prime), a, b)
	var sum Point
	sum.Add(x, x)
	println(sum.Repr())
	x = NewPoint(NewFieldElement(143, prime), NewFieldElement(98, prime), a, b)
	sum.Add(x, x)
	println(sum.Repr())
	x = NewPoint(NewFieldElement(47, prime), NewFieldElement(71, prime), a, b)
	sum.Add(x, x)
	println(sum.Repr()) // * 2
	x.Set(&sum)
	sum.Add(x, x)
	println(sum.Repr()) // * 4
	x.Set(&sum)
	sum.Add(x, x)
	println(sum.Repr()) // * 8
	origx := NewPoint(NewFieldElement(47, prime), NewFieldElement(71, prime), a, b)
	sum.Set(NewInfPoint(a, b))
	for i := 0; i < 21; i++ {
		x.Set(&sum)
		sum.Add(x, origx)
	}
	println(sum.Repr()) // *21
}

func exercise5() {
	println("Exercise 5")
	prime := int64(223)
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	origx := NewPoint(NewFieldElement(15, prime), NewFieldElement(86, prime), a, b)
	inf := NewInfPoint(a, b)
	var x Point
	x.Set(inf)
	var order int
	for i := 1; true; i++ {
		var xx Point
		xx.Set(&x)
		x.Add(&xx, origx)
		if x.Eq(inf) {
			order = i
			break
		}
	}
	println("Order of group generated:", order)
}
