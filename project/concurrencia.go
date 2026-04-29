package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	channel := make(chan string)

	for _, api := range apis {
		go checkAPI(api, channel)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-channel)
	}

	elapsed := time.Since(start)
	fmt.Printf("Time taken: %f seconds\n", elapsed.Seconds())
}

func checkAPI(api string, channel chan<- string) {
	if _, err := http.Get(api); err != nil {
		channel <- fmt.Sprintf("Error accessing API: %s, Error: %s\n", api, err.Error())
		return
	}

	channel <- fmt.Sprintf("API %s is accessible\n", api)
}
