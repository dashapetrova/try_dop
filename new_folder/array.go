package array

import (
	"bytes"
	"errors"
	"fmt"
)

// Element is a type of an array element
type Element uint64

// Array is an implementation of list using expandable array
// with fast insertion to and deletion from the end
type Array struct {
	elems []Element
	n     int
	cap   int
}

// New creates a new Array with a given capacity
func New(cap int) *Array {
	res := new(Array)
	res.n = 0
	res.cap = cap
	res.elems = make([]Element, 0, cap)

	return res
}

// Len returns the lenght of the array
func (a *Array) Len() int {
	return a.n
}

// Get retrieves an array element by index
func (a *Array) Get(i int) (Element, error) {
	if !(0 <= i && i < a.n) {
		return 0, errors.New("Index out of range")
	}

	return a.elems[i], nil
}

// Set writes an element to array by index
func (a *Array) Set(i int, x Element) error {
	if !(0 <= i && i < a.n) {
		return errors.New("Index out of range")
	}
	a.elems[i] = x

	return nil
}

// Insert adds an element to the array by index
func (a *Array) Insert(i int, x Element) error {
	if !(0 <= i && i < a.n) {
		return errors.New("Index out of range")
	}
	copy(a.elems[i+1:], a.elems[i:])
	a.elems[i] = x
	a.n = a.n + 1

	return nil
}

// Push inserts an element to the right end of the array
func (a *Array) Push(x Element) error {
	if a.cap == 0 {
		tmp_make := make([]Element, 0, 1)
		a.elems = tmp_make
		a.elems[0] = x
		a.n = a.n + 1

		return nil
	}
	if a.n == a.cap {
		tmp_make := make([]Element, a.n, a.cap*2)
		a.elems = tmp_make
	}
	a.elems[a.n] = x
	a.n = a.n + 1

	return nil
}

// Delete removes an element from the array by index
func (a *Array) Delete(i int) error {
	if !(0 <= i && i < a.n) {
		return errors.New("Index out of range")
	}
	left := a.elems[:i]
	right := a.elems[i+1:]
	a.elems = Append(a.elems, left)
	a.elems = Append(a.elems, right)
	a.n = a.n - 1

	return nil
}

// Pop deletes the last element of the array
func (a *Array) Pop() error {
	return nil
}

// String returns a textual representation of the array
func (a *Array) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i := 0; i < a.n; i += 1 {
		buffer.WriteString(a.elems[i].ToString())
		if i < a.n-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}

func (elem *Element) ToString() string {
	return fmt.Sprint(elem)
}

func Append(slice []int, elements ...int) []int {
	n := len(slice)
	total := len(slice) + len(elements)
	if total > cap(slice) {
		// Reallocate. Grow to 1.5 times the new size, so we can still grow.
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total]
	copy(slice[n:], elements)
	return slice
}
