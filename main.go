package main

import (
	"github.com/rraagg/learning-go/chapterFive"
)

func main() {
	chapterFive.DeclaringFunctions()
	chapterFive.MyFunc(chapterFive.MyFuncOpts{
		FirstName: "Barack",
		LastName:  "Obama",
		Age:       59,
	})
	chapterFive.VariadicParameters("Barack", "Michelle")

}
