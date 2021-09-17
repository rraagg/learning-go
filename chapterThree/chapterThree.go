package chapterThree

import "fmt"

func Arrays() {
	var normalArray [5]int
	normalArray[0] = 1
	fmt.Printf("Arrays are rarely used explicitly. There are limitations (you must declare the size). Array Length: %v, Capacity: %v\n", len(normalArray), cap(normalArray))
}

func Slices() {
	var normalSlice []int
	normalSlice = append(normalSlice, 1)
	fmt.Printf("Slices are normally preferred over arrays. You must use append() to grow a slice. Slice Length: %v, Capacity: %v\n", len(normalSlice), cap(normalSlice))
}

func DeclaredSlices() {
	var varDeclaration []string
	var emptySliceLiteral = []string{}
	var sliceLiteral = []string{"I'm", "a", "slice", "literal"}
	sliceOfSlices := [][]string{varDeclaration, emptySliceLiteral, sliceLiteral}

	for _, s := range sliceOfSlices {
		result := s == nil
		fmt.Printf("%v is nil: %v\n", s, result)
	}

}

func SliceOfSlice() {
	var parentSlice = []int{1, 2, 3, 4}
	var subSlice = parentSlice[:2]
	fmt.Printf("The capacity of a subslice is the same as the parent. Slice Length: %v, Capacity: %v, Subslice: %v\n", len(subSlice), cap(subSlice), subSlice)
	fmt.Printf("Slice Length: %v, Capacity: %v, Parent slice: %v\n", len(parentSlice), cap(parentSlice), parentSlice)
	parentSlice[0] = 100
	fmt.Printf("The capacity of a subslice is the same as the parent. Slice Length: %v, Capacity: %v, Subslice: %v\n", len(subSlice), cap(subSlice), subSlice)
	fmt.Printf("Slice Length: %v, Capacity: %v, Parent slice: %v\n", len(parentSlice), cap(parentSlice), parentSlice)
	subSlice = append(subSlice, 200)
	fmt.Printf("The capacity of a subslice is the same as the parent. Slice Length: %v, Capacity: %v, Subslice: %v\n", len(subSlice), cap(subSlice), subSlice)
	fmt.Printf("Slice Length: %v, Capacity: %v, Parent slice: %v\n", len(parentSlice), cap(parentSlice), parentSlice)
	fmt.Printf("Possibly never use append with subslices\nOR use a *full slice expression*")
}

func MakedSlice() {
	makedSlice := make([]int, 0, 1)
	fmt.Printf("Slice Length: %v, Capacity: %v, Maked slice: %v\n", len(makedSlice), cap(makedSlice), makedSlice)
	fmt.Printf("")
}

func CopyArrayToSlice() {
	myArray := [4]int{1, 2, 3, 4}
	mySlice := make([]int, 4)
	num := copy(mySlice, myArray[:])
	fmt.Println(mySlice, num)
}

func StringsAndRunesAndBytes() {
	str := `
	Strings are immutable.
	Strings are composed of a sequence of bytes.
	A code point in UTF can be anywhere from 1 to 4 bytes long.
	Only slice strings that you know the characters only take up one byte.
	You cannot convert an int into a string by doing a type conversion.
	A string can be converted back and forth to a slice of bytes or a slice of runes. 
	é`
	fmt.Println(str)
	fmt.Printf("I'm a slice of bytes: %v\n", len([]byte(str)))
	fmt.Printf("I'm a slice of runes: %v\n", len([]rune(str)))

}
func Utf8() {
	str := `
	UTF-8 is the most commonly used encoding for Unicode é
	- is able to expand up to UTF-32 when necessary
	- you don't have to worry about little-endian versus big-endian
	- it allows you to look at any byte in a sequence and tell if you are at the start of a UTF-8 sequence, or somewhere in the middle.
	- - This means you can't accidentally read a character incorrectly
	`
	fmt.Println(str)
	fmt.Printf("I'm a slice of bytes: %v\n", len([]byte(str)))
	fmt.Printf("I'm a slice of runes: %v\n", len([]rune(str)))

}

func DeclaringMaps() {
	var nilMap map[string]int
	fmt.Printf("Value of nilMap: %v\n", nilMap == nil)

	str := `
	Attempting to write to a nil map causes a panic
	Maps and Slices are not comparable, so they cannot be used as keys.
	The map in Go is a hash map.
	`
	fmt.Println(str)

	emptyMapLiteral := map[string]int{}
	fmt.Printf("Value of emptyMapLiteral: %v\n", emptyMapLiteral == nil)
	emptyMapLiteral["first"] = 1
	emptyMapLiteral["second"] = 2
	fmt.Printf("Value of emptyMapLiteral: %v\n", emptyMapLiteral)
	fmt.Printf(`Value of emptyMapLiteral["one"] : %v`+"\n", emptyMapLiteral["one"])

	makedMap := make(map[string]int, 5)
	fmt.Printf("Value of makedMap: %v\n", makedMap)

	makedMap["one"] = 1
	v, ok := makedMap["two"]
	fmt.Printf("Value of makedMap['two'], ok: %v, %v\n", v, ok)
}

func DeleteFromMap() {
	newMap := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Printf(`Value before delete: %v`+"\n", newMap)
	delete(newMap, "two")
	fmt.Printf(`Value after delete: %v`+"\n", newMap)
}

func MapsAsSets() {
	str := `
	Maps can be used as sets.
	You can use structs as sets to save memory usage if necessary, but maps are the preferred way. 
	`
	fmt.Println(str)

	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(intSet)
}

func Structs() {
	str := `
	Maps cannot define an API because there's no way to constrain a map to only allow certain keys.
	All the values of a map have to be the same type.
	Go doesn't have classes or inheritance.
	`
	fmt.Println(str)

	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	newPerson := Person{}

	newPerson.firstName = "Barack"
	newPerson.lastName = "Obama"
	newPerson.age = 60

}

func AnonymousStruct() {
	str := `
	Anonymous struct is a data type that is only associated with a single instance which can be used with JSON.
	Also, writing tests is another place to use anonymous structs.
	`
	fmt.Println(str)

	var person struct {
		name string
		age  int
	}
	person.name = "Barack"
	person.age = 60

	pet := struct {
		name  string
		breed string
	}{
		name:  "Jojo",
		breed: "Pyrenees",
	}
	fmt.Println(person)
	fmt.Println(pet)
}
