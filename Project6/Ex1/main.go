// Maktabkhooneh
// The project of the 6th Season - the first part

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func responseSize(url string) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	status := response.StatusCode

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(
		"Step4: ", "Body size:", len(body),
		"<=> status code:", status,
	)

}

func main() {
	// Add a count of Two one for each goroutine.
	wg.Add(2)
	fmt.Println("Start Goroutine")

	go responseSize("https://maktabkhooneh.org")
	go responseSize("https://maktabkhooneh.org/mag/")

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Terminating Program")
}
