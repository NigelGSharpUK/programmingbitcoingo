package main

import (
	. "programmingbitcoingo/code-ch01/ecc"
)

func main() {
	println("Exercise 1 is in ecc.go")
	exercise2()
	println("Exercise 3 is in ecc.go")
	exercise4()
}

func exercise2() {
	println("Exercise 2")

	prime := 57
	println((44 + 33) % prime)

	// % in Go is different to % in Python
	// We have to use our own function Mod() for -ve values (also works for +ve)
	println(Mod(9-29, prime))
	println(Mod(17+42+49, prime))
	println(Mod(52-30-38, prime))
}

func exercise4() {
	println("Exercise 4")
	prime := 97
	println(Mod(95*45*31, prime))
	println(Mod(17*13*19*44, prime))
	// Good idea to use a new PowMod() function designed to work like Python's 3 argument fn
	println(Mod(PowMod(12, 7, prime)*PowMod(77, 49, prime), prime))
}
