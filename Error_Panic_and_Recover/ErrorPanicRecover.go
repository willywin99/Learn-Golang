// // Error
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var number int
// 	var err error

// 	number, err = strconv.Atoi("123GH")

// 	if err == nil {
// 		fmt.Println(number)
// 	} else {
// 		fmt.Println(err.Error())
// 	}

// 	number, err = strconv.Atoi("123")

// 	if err == nil {
// 		fmt.Println(number)
// 	} else {
// 		fmt.Println(err.Error())
// 	}
// }

// // Error (Custom error)
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	var password string

// 	fmt.Scanln(&password)

// 	if valid, err := validPassword(password); err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println(valid)
// 	}
// }

// func validPassword(password string) (string, error) {
// 	pl := len(password)

// 	if pl < 5 {
// 		return "", errors.New("password has to have more than 4 characters")
// 	}
// 	return "Valid password", nil
// }

// // Panic
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	var password string

// 	fmt.Scanln(&password)

// 	if valid, err := validPassword(password); err != nil {
// 		panic(err.Error())
// 	} else {
// 		fmt.Println(valid)
// 	}
// }

// func validPassword(password string) (string, error) {
// 	pl := len(password)

// 	if pl < 5 {
// 		return "", errors.New("password has to have more than 4 characters")
// 	}

// 	return "Valid password", nil
// }

// Recover
package main

import (
	"errors"
	"fmt"
)

func main() {
	defer catchErr()

	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "Valid password", nil
}
