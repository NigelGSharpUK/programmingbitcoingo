package ecc

import (
	"errors"
	"strconv"
)

type FieldElement struct {
	num   int // [  ] Someday soon we'll need numbers bigger than int
	prime int
}

func NewFieldElement(num int, prime int) (*FieldElement, error) {
	if num >= prime || num < 0 {
		return nil, errors.New("Num " + strconv.Itoa(num) + " not in field range 0 to " + strconv.Itoa(prime-1))
	}
	fe := new(FieldElement)
	fe.num = num
	fe.prime = prime
	return fe, nil
}

func (fe *FieldElement) repr() string {
	return "FieldElement_" + strconv.Itoa(fe.prime) + "(" + strconv.Itoa(fe.num) + ")"
}

// Test for equality
func (fe *FieldElement) eq(other *FieldElement) bool {
	if fe == nil || other == nil {
		panic("Cannot compare nil pointers")
	}
	return fe.num == other.num && fe.prime == other.prime
}

// Test for inequality
func (fe *FieldElement) ne(other *FieldElement) bool {
	//panic("Not Implemented")

	// Answer Exercise 1
	if fe == nil || other == nil {
		panic("Cannot compare nil pointers")
	}
	return !fe.eq(other)
}

func (fe *FieldElement) add(other *FieldElement) *FieldElement {
	if fe == nil || other == nil {
		panic("Cannot add nil pointers")
	}
	if fe.prime != other.prime {
		panic("Cannot add two numbers in different Fields")
	}
	num := (fe.num + other.num) % fe.prime // [  ] Warning % only works like Python for +ve num
	res, err := NewFieldElement(num, fe.prime)
	if err != nil {
		panic(err.Error())
	}
	return res
}

// Go's % operator is DIFFERENT to Python's % operator
func mod(a, b int) int {
	return (a%b + b) % b
}

func (fe *FieldElement) sub(other *FieldElement) *FieldElement {
	//panic("Not Implemented")

	// Answer Exercise 3
	if fe == nil || other == nil {
		panic("Cannot subtract nil pointers")
	}
	if fe.prime != other.prime {
		panic("Cannot subtract two numbers in different Fields")
	}
	num := mod((fe.num - other.num), fe.prime)
	res, err := NewFieldElement(num, fe.prime)
	if err != nil {
		panic(err.Error())
	}
	return res
}

func (fe *FieldElement) mul(other *FieldElement) *FieldElement {
	//panic("Not Implemented")

	// Answer Exercise 6
	if fe == nil || other == nil {
		panic("Cannot multiply nil pointers")
	}
	if fe.prime != other.prime {
		panic("Cannot multiply two numbers in different Fields")
	}
	num := mod((fe.num * other.num), fe.prime)
	res, err := NewFieldElement(num, fe.prime)
	if err != nil {
		panic(err.Error())
	}
	return res
}

// Need a pow function with modulus, like in Python
func pow(base int, exp int, modulus int) int {
	if exp < 0 {
		panic("Negative exponent not supported")
	}
	if exp == 0 {
		return 1
	} else if exp == 1 {
		return mod(base, modulus)
	} else {
		res := 1
		for i := 0; i < exp; i++ {
			res = mod(res*base, modulus)
		}
		return res
	}
}

func (fe *FieldElement) pow(exp int) *FieldElement {
	n := mod(exp, (fe.prime - 1))
	num := pow(fe.num, n, fe.prime)
	res, err := NewFieldElement(num, fe.prime)
	if err != nil {
		panic(err.Error)
	}
	return res
}

func (fe *FieldElement) truediv(other *FieldElement) *FieldElement {
	//panic("Not Implemented")

	// Answer Exercise 9
	if fe == nil || other == nil {
		panic("Cannot divide nil pointers")
	}
	if fe.prime != other.prime {
		panic("Cannot divide two numbers in different Fields")
	}
	// Using Fermat's Little Theorem
	num := mod(fe.num*pow(other.num, fe.prime-2, fe.prime), fe.prime)
	res, err := NewFieldElement(num, fe.prime)
	if err != nil {
		panic(err.Error)
	}
	return res
}
