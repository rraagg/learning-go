package practice

import (
	"fmt"
	"strconv"
)

func StringToInt() {
	n, err := strconv.Atoi("345")
	k := strconv.Itoa(345)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v = %T\n", n, n)
	fmt.Printf("%v = %T\n", k, k)
}

func StringToAltBaseInt() {
	s := "1604"
	i, _ := strconv.ParseInt(s, 8, 0)
	fmt.Printf("%o\n", i)
}

func IntToString() {
	int := 42
	s := strconv.Itoa(int)
	fmt.Println(s)
}
