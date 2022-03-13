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
	if other == nil {
		return false
	}
	return fe.num == other.num && fe.prime == other.prime
}

// Test for inequality
func (fe *FieldElement) ne(other *FieldElement) bool {
	//panic("Not Implemented")
	return !fe.eq(other) // Answer Exercise 1
}
