package main

// A simple example from the gomock README.

// Foo is an interface with method Bar.
type Foo interface {
	Bar(x int) int
}

// aFoo defines a struct with a field that implements the Foo interface.
//
// Note, a type implements an interface by simply implementing its methods.
type aFoo struct {
	Baz Foo
}

// Baz defines a struct that implements the Foo interface.
//
// Note, a type implements an interface by simply implementing its methods.
type Baz struct{}

func (b Baz) Bar(x int) int {
	return x + 1
}

// Here, we create a new, global struct of type Baz.
var aBaz = aFoo{
	Baz: Baz{},
}

// SUT takes an argument of type Foo and calls Bar with a value.
//
// Fun fact: SUT is short for "Stuff Under Test".
func SUT(f Foo) {
	// ...
	f.Bar(99)
}

// Some more stuff under test. Here, instead of passing an argument of type
// Foo, the method uses the global struct of type Baz, which implements Foo.
//
// Looking at the test, we
func MoreSUT() {
	// ...
	aBaz.Baz.Bar(99)
}
