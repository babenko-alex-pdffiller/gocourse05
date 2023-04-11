package main

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	// Створюємо мапу для зберігання працівників
	employees := make(map[string]Employee)

	// Додаємо декілька працівників до мапи
	employees["John"] = Employee{"John", 30, 50000}
	employees["Alice"] = Employee{"Alice", 25, 40000}

	// Друк мапи перед видаленням працівника
	fmt.Println("Before deleting an employee:", employees)

	// Видаляємо працівника з мапи
	delete(employees, "John")

	// Друк мапи після видалення працівника
	fmt.Println("After deleting an employee:", employees)

	// Додаємо нового працівника до мапи
	employees["Bob"] = Employee{"Bob", 35, 60000}

	// Друк мапи після додавання нового працівника
	fmt.Println("After adding a new employee:", employees)
}
