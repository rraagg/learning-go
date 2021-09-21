package chapterFour

import "fmt"

func Blocks() {
	str := `
	Each place where a declaration occurs is called a block.
	- Package block
	- File block
	- Universe block
	 - Block that contains all other blocks
	 - Go only contains 25 key words, but built-in types aren't included in that list.
	 - Predeclared identifiers are defined in the Universe Block
	 - BE VERY CAREFUL TO NOT REDEFINE PREDECLARED IDENTIFIERS 
	`
	fmt.Println(str)
}

func ShadowingVariables() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

func IfStatement() {
	str := `
	Variables declared within the braces of an if statement are only available in that block.
	You can declare variables that are scoped to the condition AND the if and else blocks.
	Go encourages short if statement bodies, left-aligned as much as possible
	`
	fmt.Println(str)
}

func ForLoop() {
	str := `
	For is the only looping mechanism
	- A complete, C-style for
	- A condition-only for
	- An infinite for
	 - break
	 - continue
	- for-range


	`
	fmt.Println(str)
}

func ForRange() {
	str := `
	You can only use the for-range loop to iterate over the built-in compound types
	and the user-defined types that are based on them.
	Use a for-range loop to access the runes in a sting in order. The key is the number of bytes
	from the beginning of the string, but the type of the value is rune.
	`
	fmt.Println(str)
}

func WhichForLoop() {
	str := `
	Most times you will use a for-range loop, unless you need to start at a certain index.
	Standard loop does not properly handle multi-byte characters.
	`
	fmt.Println(str)
}

func SwitchStatement() {
	str := `
By default, cases in switch statements don't fall through'
In Go, An empty case means nothing happens.
Blank switches do not specify the value that you are comparing against.
	`
	fmt.Println(str)

	words := []string{
		"a",
		"cow",
		"smile",
		"gopher",
		"octopus",
		"anthopologist",
	}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word")
		}
	}
}

func BlankSwitches() {

	str := `
By default, cases in switch statements don't fall through'
In Go, An empty case means nothing happens.
Blank switches do not specify the value that you are comparing against.
	`
	fmt.Println(str)

	words := []string{
		"hi",
		"salutations",
		"hello",
	}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word")
		case wordLen > 10:
			fmt.Println(word, "is a long word")
		default:
			fmt.Println(word, "is exactly the right length")
		}
	}
}
