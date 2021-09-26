package chapterFive

import (
	"fmt"
	"time"
)

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
	Address   string
}

type CurrentOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func DeclaringFunctions() {
	str := `
	Go does not have named or optional input parameters.
	If you want to emulate this behavior you can define a struct
	and then pass the struct into the function
	`
	fmt.Println(str)
	fmt.Println("Function Declared!")

}

func MyFunc(opts MyFuncOpts) {
	fmt.Println(opts.FirstName)
}

func VariadicParameters(names ...string) {

	for _, name := range names {
		fmt.Println(name)

	}
}

func MultipleReturnValues(opts MyFuncOpts, ln string) (fn string, age int) {
	str := `
	Python returns multiple values in a Tuple and you can just assign the
	results of the function to one variable.
	In Go, you must assign both values returned to separate variables. 
	`
	fmt.Println(str)
	if opts.LastName == ln {
		return opts.FirstName, opts.Age
	}
	return "no match", 0
}

func BirthDateValidator() {
	birthDate := time.Date(1901, time.January, 1, 1, 0, 0, 0, time.Local)
	baseDate := time.Date(1900, time.January, 1, 1, 0, 0, 0, time.Local)
	if birthDate.Year() < baseDate.Year() {
		fmt.Println("Please enter a birth date with a year greater than 1900")
	} else {
		fmt.Println("your date is fine")
	}

}

func CallByValue(i int, s string, m MyFuncOpts) {
	i = i * 2
	s = "Michelle"
	m.FirstName = "Michelle"

}
