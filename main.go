package main

import (
	"fmt"

	"github.com/rraagg/learning-go/chapterFive"
)

func main() {
	chapterFive.DeclaringFunctions()
	myfuncOpts := chapterFive.MyFuncOpts{
		FirstName: "Barack",
		LastName:  "Obama",
		Age:       59,
		Address:   "Chicago",
	}
	currentOpts := chapterFive.CurrentOpts{
		FirstName: "Barack",
		LastName:  "Obama",
		Age:       59,
	}

	chapterFive.VariadicParameters("Barack", "Michelle")
	optsID := make(map[chapterFive.CurrentOpts]chapterFive.MyFuncOpts)
	optsID[currentOpts] = myfuncOpts
	currentOpts.FirstName = "Michelle"

	fmt.Println(currentOpts)
	fmt.Println(optsID[currentOpts])

	first, age := chapterFive.MultipleReturnValues(myfuncOpts, "Obam")
	fmt.Println(first, age)

	chapterFive.BirthDateValidator()
	i := 10
	s := "No change"
	m := myfuncOpts
	chapterFive.CallByValue(i, s, m)
	fmt.Println(i, s, m)
}
