// // Channels (Channel operators)
// func main() {
// 	c := make(chan string)

// 	// Mengirim data kepada channel
// 	c <- value

// 	// Menerima data dari channel
// 	result := <-c
// }

// // Channels (Implementing Channel)
// package main

// import "fmt"

// func main() {
// 	c := make(chan string)
// 	go introduce("Airell", c)
// 	go introduce("Nanda", c)
// 	go introduce("Mailo", c)

// 	msg1 := <-c
// 	fmt.Println(msg1)

// 	msg2 := <-c
// 	fmt.Println(msg2)

// 	msg3 := <-c
// 	fmt.Println(msg3)

// 	close(c)
// }

// func introduce(student string, c chan string) {
// 	result := fmt.Sprintf("Hai, my name is %s", student)
// 	c <- result
// }

// // Channels (Channel with anonymous function)
// package main

// import "fmt"

// func main() {
// 	c := make(chan string)
// 	students := []string{"Airell", "Mailo", "Indah"}

// 	for _, v := range students {
// 		go func(student string) {
// 			fmt.Println("Student", student)
// 			result := fmt.Sprintf("Hai, my name is %s", student)
// 			c <- result
// 		}(v)
// 	}

// 	for i := 1; i < 3; i++ {
// 		print(c)
// 	}
// 	close(c)
// }

// func print(c chan string) {
// 	fmt.Println(<-c)
// }

// // Channels (Channel directions)
// package main

// import "fmt"

// func main() {
// 	c := make(chan string)
// 	students := []string{"Airell", "Mailo", "Indah"}

// 	for _, v := range students {
// 		go introduce(v, c)
// 	}

// 	for i := 1; i < 3; i++ {
// 		print(c)
// 	}

// 	close(c)
// }

// func print(c <-chan string) {
// 	fmt.Println(<-c)
// }

// func introduce(student string, c chan<- string) {
// 	result := fmt.Sprintf("Hai, my name is %s", student)
// 	c <- result
// }

// // UnBuffered Channel (Buffered vs UnBuffered Channel)
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c1 := make(chan int)

// 	go func(c chan int) {
// 		fmt.Println("func goroutine starts sending data into the channel")
// 		c <- 10
// 		fmt.Println("func goroutine after sending data into channel")
// 	}(c1)

// 	fmt.Println("main goroutine sleep for 2 seconds")
// 	time.Sleep(time.Second * 2)

// 	fmt.Println("main goroutine starts receiving data")
// 	d := <-c1
// 	fmt.Println("main goroutine received data:", d)

// 	close(c1)
// 	time.Sleep(time.Second)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c1 := make(chan int, 3)

// 	go func(c chan int) {
// 		for i := 1; i <= 5; i++ {
// 			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
// 			c <- i
// 			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
// 		}
// 		close(c)
// 	}(c1)

// 	fmt.Println("main goroutine sleeps 2 seconds")
// 	time.Sleep(time.Second * 2)

// 	for v := range c1 { // v = <- c1
// 		fmt.Println("main goroutine received value from channel:", v)
// 	}
// }

// Channel (Select)
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "Salut!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}
