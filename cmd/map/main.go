package main

import (
	"fmt"
)

func main() {
	// 1 definition
	filled := map[string]int{"C": 1972, "Go": 2009}
	fmt.Printf("%T, %v\n", filled, filled)

	var none map[string]int
	fmt.Printf("%T, %v\n", none, none)

	ten := make(map[string]int, 10)
	fmt.Printf("%T, %v\n", ten, ten)

	zero := *new(map[string]int)
	fmt.Printf("%T, %v\n", zero, zero)

	//// unaddressable
	//languagesYears := map[string]int{"C": 1972, "Go": 2009}
	//languages := []string{"C", "Go"}
	//_ = &languagesYears[1]
	//_ = &languages[`C`]

	// incomparable
	//incomparable1 := map[[]int]string{}
	//incomparable2 := map[map[string]any]string{}
	//incomparable3 := map[func() bool]string{}

	// 2 init and update
	type person struct{ age, height int }
	persons := map[string]person{}
	persons["John"] = person{age: 29}
	fmt.Printf("%v\n", persons)
	delete(persons, "John")

	numbers := map[int][5]int{}
	numbers[1] = [5]int{1: 789}
	fmt.Printf("%v\n", numbers)
	delete(numbers, 1)

	// exercise // writing
	// myFirstMap := map[string]string{
	// 	`mom name`: `Name`,
	// 	`dad name`: `Name`,
	// 	`pet name`: `Name`,
	// }
	// fmt.Printf("%T, %v\n", myFirstMap, myFirstMap)

	// exercise  // read it
	// pm1_1 := map[[2]int]int{{1, 2}: 1972, {3, 4}: 2009}
	// pm1_2 := map[[2]int]int{[2]int{1, 2}: 1972, [2]int{3, 4}: 2009}
	// pm2 := map[bool]string{} // why is bad?

	// exercise: update age
	// relatives := make(map[string]int, 3)
	// relatives[`mom name`] = 50
	// relatives[`dad name`] = 60
	// relatives[`pet name`] = 10
	// fmt.Printf("%T, %v\n", relatives, relatives)
}
