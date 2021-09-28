package chapterSix

import "fmt"

func Pointers() {
	str := `
A pointer is a variable that stores the memory location of a value.
* is the indirection operator. It precedes the variable of pointer type
and returns the value pointed at.
& is the address operator. It precedes a value type and returns the address
of the memory location where the value is stored.
You must check if a pointer is nil before dereferencing it. A nil pointer
will cause the program to panic if a dereference is attempted on it.
`
	fmt.Println(str)
}
