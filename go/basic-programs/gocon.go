package main 

import( "fmt"
		"time"
		)
func main() {
	fmt.Println("GO concurrency")

	go func() {
		fmt.Println("Insode goroutine")
	}()

	time.Sleep(100000) // Wait is necessary

}