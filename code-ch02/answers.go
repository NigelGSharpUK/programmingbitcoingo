package main

import (
	. "programmingbitcoingo/code-ch02/ecc"
)

func main() {
	exercise1()
}

func exercise1() {
	p, err := NewPointErr(2, 4, 5, 7)
	if err != nil {
		println("Point is NOT on the curve: ", 2, ", ", 4)
	} else {
		println("Point is on the curve: ", p.Repr())
	}

	p, err = NewPointErr(-1, -1, 5, 7)
	if err != nil {
		println("Point is NOT on the curve: ", -1, ", ", -1)
	} else {
		println("Point is on the curve: ", p.Repr())
	}

	p, err = NewPointErr(18, 77, 5, 7)
	if err != nil {
		println("Point is NOT on the curve: ", 18, ", ", 77)
	} else {
		println("Point is on the curve: ", p.Repr())
	}

	p, err = NewPointErr(5, 7, 5, 7)
	if err != nil {
		println("Point is NOT on the curve: ", 5, ", ", 7)
	} else {
		println("Point is on the curve: ", p.Repr())
	}
}
