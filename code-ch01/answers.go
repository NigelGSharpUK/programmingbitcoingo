package main

import (
	"programmingbitcoingo/code-ch01/ecc"
)

func main() {
	println("Exercise 1 is in ecc.go")
	exercise2()
}

func exercise2() {
	println("Exercise 2")

	prime := 57
	println((44 + 33) % prime)

	// % in Go is different to % in Python
	// We have to use our own function Mod() for -ve values
	println(ecc.Mod(9-29, prime))
	println(ecc.Mod(17+42+49, prime))
	println(ecc.Mod(52-30-38, prime))
}
