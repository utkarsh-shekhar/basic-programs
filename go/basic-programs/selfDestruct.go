package main

import "fmt"
import "time"

func cancel(abort chan bool) {
	fmt.Printf("This program will self destruct, doyou wish to cancel?\n")
var r int
fmt.Scanf("%c", &r)
switch r {
	default: abort <- false
	case 'Y': abort <- true
	case 'y': abort <- true
	}
}


func countDown(count chan int) {
	for i := 10 ; i >= 0 ; i--{
	count <- i
	time.Sleep(1000000000)
	}
}

func main() {
	abort := make(chan bool)
	count := make(chan int)
	go cancel(abort)
	go countDown(count)
	for {
	select {
		case i := <- countDown:
			if 0 == i {
				selfDestruct()
				return
			}
		fmt.Printf("%d seconds remaining\n", i)
			case a := <- abort:
				if a {
					fmt.Printf("Self destruct aborted\n")
				} else {
					selfDestruct()
				}
			return
		}
	}
}