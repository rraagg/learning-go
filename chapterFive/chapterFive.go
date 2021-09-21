package chapterFive

import "fmt"

type MyFuncOpts struct {
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
