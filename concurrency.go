// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup

// func main() {

// 	wg.Add(3)

// 	go greetings("Hello World")
// 	go greetings("Hello Chandrashekhar")
// 	go greetings("Hello Sham")

// 	wg.Wait()

// 	fmt.Println("Hello Its Done...")

// }

// func greetings(s string) {

// 	defer wg.Done()

// 	runes := []rune(s)

// 	for i := 0; i < len(s); i++ {

// 		fmt.Print(string(runes[i]))
// 		time.Sleep(100 * time.Millisecond)

// 	}
// 	fmt.Println()
// }
