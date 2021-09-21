package main

import (
	"github.com/rraagg/learning-go/chapterFive"
)

func main() {
	chapterFive.DeclaringFunctions()
	chapterFive.MyFunc(chapterFive.MyFuncOpts{
		FirstName: "Rob",
		LastName:  "Gallagher",
		Age:       44,
	})
}
