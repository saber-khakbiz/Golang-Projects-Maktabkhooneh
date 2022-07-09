// Golang program to read and write the files
package main

// importing the packages
import (
	"fmt"
	"time"
	"io/ioutil"
	"log"
	"os"
)

var path = "test.txt"
func WriteFile(massage string) {


	fmt.Printf("Writing to a file in Go lang\n")
	

	file, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	
	if err != nil {
		panic(err)
	}
	

	defer file.Close()
	
	dt := time.Now().Format("Mon Jan 02 15:04:05 -0700 2006")
	dt = string(dt)

	massage = massage + " --> (" + dt + ")\n"

	len, err := file.WriteString(massage)

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}


	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile() {

	fmt.Printf("\n\nReading a file in Go lang\n")
	fileName := "test.txt"
	
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

// main function
func main() {
	massage := "hello, add a new line to this file!"
	WriteFile(massage)
	ReadFile()
}
