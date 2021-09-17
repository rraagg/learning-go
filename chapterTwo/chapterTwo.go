package chapterTwo

import "fmt"

// ChapterTwo prints a zero value
func ZeroValue() {
	var zeroValue int
	fmt.Println(zeroValue)
}

func IntegerLiteral() {
	var integerLiteral = 0o16
	newIntegerLiteral := integerLiteral + 0o14
	fmt.Println(newIntegerLiteral)
}

func FloatingPointLiteral() {
	var floatingPointLiteral = 6.03e23
	fmt.Println(floatingPointLiteral)
}

func RuneLiteral() {
	var runeLiteral = 'a'
	fmt.Println(runeLiteral)
}

func StringLiteral() {
	var stringLiteral = "Greetings and Salutations"
	fmt.Println(stringLiteral)
}

func RawStringLiteral() {
	var rawStringLiteral = `
	Greetings and
	Salutations
	`
	fmt.Println(rawStringLiteral)
}

func BooleanLiteral() {
	var booleanLiteral = true
	fmt.Println(booleanLiteral)
	fmt.Printf("%T", booleanLiteral)
}

func IntegerTypes() {
	var iByte byte
	var iUint8 uint8
	iByte = 1
	iUint8 = 2
	total := iByte + iUint8
	fmt.Printf("Add byte and uint8: %v\n", total)

	var iInt int
	var iInt32 int32
	iInt = 1
	iInt32 = 2
	total = byte(iInt) + byte(iInt32)
	fmt.Printf("Type conversion needed to add int and int32: %v\n", total)
}

func DeclarationList() {
	var (
		x int    = 1
		y string = "One"
	)
	fmt.Printf("Declaration List: %v and %s\n", x, y)
}
