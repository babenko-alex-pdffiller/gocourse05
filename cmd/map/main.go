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

	// exercises
	// 1. Створити словник англійських слів з їх українським перекладом та знайти переклад певного слова.
	// 3. Створити програму, яка буде зберігати дані про працівників певної компанії, додапє і видаляє працівників
}
