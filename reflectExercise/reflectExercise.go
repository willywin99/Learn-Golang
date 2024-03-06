package main

import (
	"fmt"
	"reflect"
)

// // Identifying Data Type & Value
// func main() {
// 	var number = 23
// 	var reflectValue = reflect.ValueOf(number)

// 	fmt.Println("tipe variabel :", reflectValue.Type())

// 	if reflectValue.Kind() == reflect.Int {
// 		fmt.Println("nilai variabel :", reflectValue.Int())
// 	}
// }

// // Accessing Value using Interface Method
// func main() {
// 	var number = 23
// 	var reflectValue = reflect.ValueOf(number)
// 	fmt.Println("tipe variabel :", reflectValue.Type())
// 	fmt.Println("nilai variabel :", reflectValue.Interface())

// 	var nilai = reflectValue.Interface().(int)
// 	fmt.Println(nilai)
// }

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var s1 = &student{Name: "john wick", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})
	fmt.Println("nama :", s1.Name)
}
