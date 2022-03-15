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
	//return y.Pow(2).Eq(x.Pow(3).Add(a.Mul(x)).Add(b))
	return y.Pow(big.NewInt(2)).Eq(x.Pow(big.NewInt(3)).Add(a.Mul(x)).Add(b))
}

func exercise1() {
	println("Exercise 1")
	prime := big.NewInt(223)
	a := NewFieldElement(big.NewInt(0), prime)
	b := NewFieldElement(big.NewInt(7), prime)
	println(on_curve(NewFieldElement(big.NewInt(192), prime), NewFieldElement(big.NewInt(105), prime), a, b))
	println(on_curve(NewFieldElement(big.NewInt(17), prime), NewFieldElement(big.NewInt(56), prime), a, b))
	println(on_curve(NewFieldElement(big.NewInt(200), prime), NewFieldElement(big.NewInt(119), prime), a, b))
	println(on_curve(NewFieldElement(big.NewInt(1), prime), NewFieldElement(big.NewInt(193), prime), a, b))
	println(on_curve(NewFieldElement(big.NewInt(42), prime), NewFieldElement(big.NewInt(99), prime), a, b))
}

func exercise2() {
	println("Exercise 2")
	prime := big.NewInt(223)
	a := NewFieldElement(big.NewInt(0), prime)
	b := NewFieldElement(big.NewInt(7), prime)
	x := NewPoint(NewFieldElement(big.NewInt(170), prime), NewFieldElement(big.NewInt(142), prime), a, b)
	y := NewPoint(NewFieldElement(big.NewInt(60), prime), NewFieldElement(big.NewInt(139), prime), a, b)
	println(x.Add(y).Repr())
	x = NewPoint(NewFieldElement(big.NewInt(47), prime), NewFieldElement(big.NewInt(71), prime), a, b)
	y = NewPoint(NewFieldElement(big.NewInt(17), prime), NewFieldElement(big.NewInt(56), prime), a, b)
	println(x.Add(y).Repr())
	x = NewPoint(NewFieldElement(big.NewInt(143), prime), NewFieldElement(big.NewInt(98), prime), a, b)
	y = NewPoint(NewFieldElement(big.NewInt(76), prime), NewFieldElement(big.NewInt(66), prime), a, b)
	println(x.Add(y).Repr())
}

func exercise4() {
	println("Exercise 4")
	prime := big.NewInt(223)
	a := NewFieldElement(big.NewInt(0), prime)
	b := NewFieldElement(big.NewInt(7), prime)
	x := NewPoint(NewFieldElement(big.NewInt(192), prime), NewFieldElement(big.NewInt(105), prime), a, b)
	println(x.Add(x).Repr())
	x = NewPoint(NewFieldElement(big.NewInt(143), prime), NewFieldElement(big.NewInt(98), prime), a, b)
	println(x.Add(x).Repr())
	x = NewPoint(NewFieldElement(big.NewInt(47), prime), NewFieldElement(big.NewInt(71), prime), a, b)
	x = x.Add(x)
	println(x.Repr()) // *2
	x = x.Add(x)
	println(x.Repr()) // *4
	x = x.Add(x)
	println(x.Repr()) // *8
	origx := NewPoint(NewFieldElement(big.NewInt(47), prime), NewFieldElement(big.NewInt(71), prime), a, b)
	x = NewInfPoint(a, b)
	for i := 0; i < 21; i++ {
		x = x.Add(origx)
	}
	println(x.Repr()) // *21
}

func exercise5() {
	println("Exercise 5")
	prime := big.NewInt(223)
	a := NewFieldElement(big.NewInt(0), prime)
	b := NewFieldElement(big.NewInt(7), prime)
	origx := NewPoint(NewFieldElement(big.NewInt(15), prime), NewFieldElement(big.NewInt(86), prime), a, b)
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
