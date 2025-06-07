package practice

import (
	"fmt"
	"strconv"
)

func StringToInt() {
	n, err := strconv.Atoi("345")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v = %T\n", n, n)
}
