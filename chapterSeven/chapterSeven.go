package chapterSeven

import "fmt"

type Country struct {
	name           string
	continent      string
	population     uint64
	nationalAnimal string
}

func (c Country) GetContinent() string {
	return c.continent
}

func Methods() {
	fmt.Println("vim-go")
}
