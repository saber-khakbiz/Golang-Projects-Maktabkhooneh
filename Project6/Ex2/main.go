// Maktabkhooneh
// The project of the 6th Season - the second part

package main

import (
	"fmt"
)

func strlen(txt string, c chan<- int) {
	c <- len(txt)
}

func main() {
	c := make(chan int)
	go strlen("Hi, my name is saber.", c)
	go strlen("Maktabkhooneh", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}
