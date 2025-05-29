package chapter_thirteen

import "time"

type Person struct {
	Name      string
	Age       int
	DateAdded time.Time
}

func CreatePerson(name string, age int) Person {
	return Person{
		Name:      name,
		Age:       age,
		DateAdded: time.Now(),
	}
}

func AddNumbers(x, y int) int {
	return x + y
}

func addNumbers(x, y int) int {
	return x + y
}
