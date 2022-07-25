// Maktabkhooneh
// The project of the 6th Season - the second part

package main

import (
	"fmt"
)

// AddHello Add {Hello} word  to text
func AddHello(ch chan<- string, txt string) {
	txt = "Hello " + txt
	ch <- txt
}

// LenStr Get string length
func LenStr(ch <-chan string, resch chan<- int) {
	str := <-ch
	fmt.Println(str)
	LOS := len(str)
	resch <- LOS
}

func main() {
	//Create two channel
	ch := make(chan string)
	resch := make(chan int)

	// goroutine
	go AddHello(ch, "Maktabkhooneh!")
	go LenStr(ch, resch)

	//Print OutPut of resch Channel
	result := <-resch
	fmt.Println("Result:", result, "Characters")
}
