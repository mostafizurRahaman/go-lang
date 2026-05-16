package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkUrlStatus(url string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(" 🔃 Starting checking url ⏳", url)

	// ! waiting for one second:
	time.Sleep(time.Second * 1)

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("❌ Having some issues ->", url)
	}

	fmt.Printf(" ✅ Successfully tested the url %v \n", url)
	fmt.Println(res.Status)
	fmt.Println(res.Body)

}

func main() {

	// ! Waiting Group create:
	var wg sync.WaitGroup

	// ! urls slices:
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.golang.org",
	}

	// ! Run for each  to test each one with go routine and waiting group:

	for _, url := range urls {
		// ! Add one into wait group each loop
		wg.Add(1)
		go checkUrlStatus(url, &wg)
	}

	// ! Wait for end:
	wg.Wait()

}
