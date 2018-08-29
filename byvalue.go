package main

import (
	"fmt"
	"unsafe"
)

type model struct {
	i int
}

func (m model) string() {
	fmt.Println("i = ", m.i)
}

var i int

func newModel() *model {
	i++
	m := new(model)
	m.i = i
	return m
}

// Go 是传值的
func foo() {
	a := 1
	fmt.Printf("a before %v \n", unsafe.Pointer(&a))

	b := a
	fmt.Printf("b %v \n", unsafe.Pointer(&b))

	a = 2
	fmt.Printf("a after %v \n", unsafe.Pointer(&a))

	fmt.Printf("a = %d, b = %d \n", a, b) // a = 2, b = 1
}

func bar() {
	m := model{i: 1}
	fmt.Printf("m before %v \n", unsafe.Pointer(&m))

	defer m.string() // 1

	defer func() { m.string() }() // 3

	defer func(m model) { m.string() }(m) // 1

	defer func(m *model) { m.string() }(&m) // 3

	m.i = 2
	m.string() // 2

	defer m.string() // 2

	defer func(m model) { m.string() }(m) // 2

	m = model{i: 3}
	fmt.Printf("m after %v \n", unsafe.Pointer(&m))
}

func baz() {
	m := newModel()
	fmt.Printf("m before %v \n", unsafe.Pointer(m))
	fmt.Printf("m before %v \n", unsafe.Pointer(&m))

	m.string()                               // 1
	defer m.string()                         // 1
	defer func() { m.string() }()            // 2
	defer func(_m *model) { _m.string() }(m) // 100

	m.i = 100

	m = newModel()
	m.string() // 2
	fmt.Printf("m after %v \n", unsafe.Pointer(m))
	fmt.Printf("m after %v \n", unsafe.Pointer(&m))
}

func main() {
	foo()
	bar()
	baz()
}
