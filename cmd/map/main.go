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

	// exercise // writing
	// myFirstMap := map[string]string{
	// 	`mom name`: `Name`,
	// 	`dad name`: `Name`,
	// 	`pet name`: `Name`,
	// }
	// fmt.Printf("%T, %v\n", myFirstMap, myFirstMap)

	//// unaddressable
	//languagesYears := map[string]int{"C": 1972, "Go": 2009}
	//languages := []string{"C", "Go"}
	//_ = &languagesYears[1]
	//_ = &languages[`C`]

	// exercise  // read it
	// pm1_1 := map[[2]int]int{{1, 2}: 1972, {3, 4}: 2009}
	// pm1_2 := map[[2]int]int{[2]int{1, 2}: 1972, [2]int{3, 4}: 2009}
	// pm2 := map[bool]string{} // why is bad?

	// incomparable
	//incomparable1 := map[[]int]string{}
	//incomparable2 := map[map[string]any]string{}
	//incomparable3 := map[func() bool]string{}

	// 2 init and update
	type person struct{ age, height int }
	persons := map[string]person{}
	persons["John"] = person{age: 29}
	fmt.Printf("%v\n", persons)

	numbers := map[int][5]int{}
	numbers[1] = [5]int{1: 789}
	fmt.Printf("%v\n", numbers)

	// exercise: update age
	// relatives := make(map[string]int, 3)
	// relatives[`mom name`] = 50
	// relatives[`dad name`] = 60
	// relatives[`pet name`] = 10
	// fmt.Printf("%T, %v\n", relatives, relatives)

	// 3 // read it (різниця між типом і змінною) (Як можна спростити запис знизу)
	langs := map[struct{ dynamic, strong bool }]map[string]int{
		{true, false}:  {"JavaScript": 1995},
		{false, true}:  {"Go": 2009},
		{false, false}: {"C": 1972},
	}
	destination1 := map[*struct{ dynamic, strong bool }]*map[string]int{}
	for category, langInfo := range langs {
		destination1[&category] = &langInfo // check if it will be effects on langInfo (value)
		category.dynamic, category.strong = true, true
	}
	for category, langInfo := range langs {
		fmt.Println(category, langInfo)
	}

	destination2 := map[struct{ dynamic, strong bool }]map[string]int{}
	for category, langInfo := range destination1 {
		destination2[*category] = *langInfo
	}
	// destination1 and destination2 both contain only one entry.
	fmt.Println(len(destination1), len(destination2)) // 1 1
	fmt.Println(destination2)                         // map[{true true}:map[C:1972]]
}
