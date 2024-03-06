package main

import "fmt"

// // Empty Interface
// func main() {
// 	var randomValues interface{}
// 	_ = randomValues
// 	randomValues = "Jalan Sudirman"
// 	randomValues = 20
// 	randomValues = true
// 	randomValues = []string{"Airell", "Nanda"}
// }

// // Empty Interface (Type Assertion)
// func main() {
// 	var v interface{}
// 	v = 20
// 	// v = v * 9

// 	// Empty Interface (Type Assertion)
// 	if value, ok := v.(int); ok == true {
// 		v = value * 9
// 	}
// 	fmt.Println(v)
// }

// Empty Interface (Map & Slice with Empty Interface)
func main() {
	rs := []interface{}{1, "Airell", true, 2, "Ananda", true}
	rm := map[string]interface{}{
		"Name":   "Airell",
		"Status": true,
		"Age":    23,
	}
	_, _ = rs, rm

	fmt.Println(rs)
	fmt.Println(rm)
}
