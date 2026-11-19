// package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wtobj sync.WaitGroup
// var once sync.Once

// func main() {

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	wtobj.Add(3)

// 	go Gret(ctx, "AAAA")
// 	go Gret(ctx, "BBBB")
// 	go Gret(ctx, "CCCC")

// 	wtobj.Wait()

// }

// func Gret(ctx context.Context, s string) {

// 	runes := []rune(s)
// 	defer wtobj.Done()

// 	for i := 0; i < len(runes); i++ {

// 		select {
// 		case <-ctx.Done():
// 			//sync.Once = Execute a function exactly once, regardless of how many goroutines call it.
// 			once.Do(func() {
// 				fmt.Println("Times Up")

// 			})
// 			return

// 		case <-time.After(3 * time.Second):
// 			fmt.Println(string(runes[i]))

// 		}

// 	}

// }
