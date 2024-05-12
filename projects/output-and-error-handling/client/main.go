package main

import (
	"fmt"
	"net/http"
	"time"
	"bufio"
)

func main() {
	maxRetries := 3
	retryInterval := time.Second * 2

	for retry := 0; retry < maxRetries; retry++ {
		res, err := http.Get("http://localhost:8080")
		
		if err != nil {
			if res != nil {
				fmt.Println("Response status:", res.Status)
				if res.StatusCode == http.StatusTooManyRequests {
					fmt.Println("Server is too busy. Retry after:", res.Header.Get("Retry-After"))

					retryAfterStr := res.Header.Get("Retry-After")
					retryAfter, err := time.ParseDuration(retryAfterStr)
					if err != nil {
						
						retryAfter = retryInterval
					}
					time.Sleep(retryAfter)
					continue 
				} else {
					fmt.Println("Failed to get response from server:", err)
				}
			} else {
				fmt.Println("Failed to connect to server:", err)
			}

			
			time.Sleep(retryInterval)
			continue 
		}

		defer res.Body.Close()

		fmt.Println("Response status:", res.Status)

		if res.StatusCode == http.StatusOK {
			scanner := bufio.NewScanner(res.Body)
			for i := 0; scanner.Scan(); i++ {
				fmt.Println(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				panic(err)
			}
			break 
		} else {
			fmt.Println("Error response from server:", res.Status)
			
			time.Sleep(retryInterval)
			continue 
		}
	}

	fmt.Println("Out of retries. Exiting...")
}
