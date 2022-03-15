package main

import (
	. "programmingbitcoingo/code-ch03/ecc"
)

func main() {
	exercise1()
}

func on_curve(x, y, a, b *FieldElement) bool {
	return y.Pow(2).Eq(x.Pow(3).Add(a.Mul(x)).Add(b))
}

func exercise1() {
	println("Exercise 1")
	const prime = 223
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	println(on_curve(NewFieldElement(192, prime), NewFieldElement(105, prime), a, b))
	println(on_curve(NewFieldElement(17, prime), NewFieldElement(56, prime), a, b))
	println(on_curve(NewFieldElement(200, prime), NewFieldElement(119, prime), a, b))
	println(on_curve(NewFieldElement(1, prime), NewFieldElement(193, prime), a, b))
	println(on_curve(NewFieldElement(42, prime), NewFieldElement(99, prime), a, b))
}
