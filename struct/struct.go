package main

import "fmt"

type Person struct {
	name string
	age  int
}

// type Employee struct {
// 	division string
// 	person   Person
// }

// type Employee struct {
// 	name     string
// 	age      int
// 	division string
// }

// // Struct (Giving Value to Struct)
// func main() {
// 	var employee Employee
// 	employee.name = "Airell"
// 	employee.age = 23
// 	employee.division = "Curriculum Developer"
// 	fmt.Println(employee.name)
// 	fmt.Println(employee.age)
// 	fmt.Println(employee.division)
// }

// // Struct (Initializing Struct)
// func main() {
// 	var employee1 = Employee{}
// 	employee1.name = "Airell"
// 	employee1.age = 23
// 	employee1.division = "Curriculum Developer"

// 	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}
// 	fmt.Printf("Employee1: %+v\n", employee1)
// 	fmt.Printf("Employee2: %+v\n", employee2)
// }

// // Struct (Pointer to a Struct)
// func main() {
// 	var employee1 = Employee{name: "Airell", age: 23, division: "Curriculum Developer"}
// 	var employee2 *Employee = &employee1

// 	fmt.Println("Employee1 name:", employee1.name)
// 	fmt.Println("Employee2 name:", employee2.name)

// 	fmt.Println(strings.Repeat("#", 30))

// 	employee2.name = "Ananda"

// 	fmt.Println("Employee1 name:", employee1.name)
// 	fmt.Println("Employee2 name:", employee2.name)
// }

// // Struct (Embedded Struct)
// func main() {
// 	var employee1 = Employee{}

// 	employee1.person.name = "Airell"
// 	employee1.person.age = 23
// 	employee1.division = "Curriculum Developer"

// 	fmt.Printf("%+v", employee1)
// }

// // Struct (Anonymous struct)
// func main() {
// 	// Anonymous struct tanpa pengisian field
// 	var employee1 = struct {
// 		person   Person
// 		division string
// 	}{}
// 	employee1.person = Person{name: "Airell", age: 23}
// 	employee1.division = "Curriculum developer"

// 	// Anonymous struct dengan pengisian field
// 	var employee2 = struct {
// 		person   Person
// 		division string
// 	}{
// 		person:   Person{name: "Ananda", age: 23},
// 		division: "Finance",
// 	}

// 	fmt.Printf("Employee1: %+v\n", employee1)
// 	fmt.Printf("Employee2: %+v\n", employee2)
// }

// // Struct (Slice of Struct)
// func main() {
// 	var people = []Person{
// 		{name: "Airell", age: 23},
// 		{name: "Ananda", age: 23},
// 		{name: "Mailo", age: 23},
// 	}

// 	for _, v := range people {
// 		fmt.Printf("%+v\n", v)
// 	}
// }

// Struct (Slice of Anonymous Struct)
func main() {
	var employee = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Airell", age: 23}, division: "Curriculum Developer"},
		{person: Person{name: "Ananda", age: 23}, division: "Finance"},
		{person: Person{name: "Mailo", age: 23}, division: "Marketing"},
	}

	for _, v := range employee {
		fmt.Printf("%+v\n", v)
	}
}
