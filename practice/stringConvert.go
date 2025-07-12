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

func FormatInt() {
	x := int64(345)
	s := strconv.FormatInt(x, 10)
	fmt.Printf("%T, %v\n", s, s)

	b := int64(-1024)
	t := strconv.FormatInt(b, 2)
	fmt.Printf("%T, %v\n", t, t)
}

func AppendInt() {
	b := []byte("What are we appending?: ")
	b = strconv.AppendInt(b, 22, 10)
	fmt.Printf("%v\n", string(b))
}

func ParseFloat() {
	n := "3.1415926535"
	f, err := strconv.ParseFloat(n, 64)
	fmt.Printf("%v - %T - %v\n", f, f, err)
}
