package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	maxRetries := 3
	retryInterval := time.Second * 2

	for retry := 0; retry < maxRetries; retry++ {
		res, err := http.Get("http://localhost:8080")
		if err != nil {
			fmt.Println("Failed to get response from server:", err)
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
		} else if res.StatusCode == http.StatusTooManyRequests {
			fmt.Println("Server is too busy. Retry after:", res.Header.Get("Retry-After"))

			retryAfterStr := res.Header.Get("Retry-After")
			retryAfter, err := strconv.Atoi(retryAfterStr)
			if err != nil {
				fmt.Println("Invalid Retry-After header value, defaulting to retry interval:", err)
				retryAfter = int(retryInterval.Seconds())
			}
			time.Sleep(time.Duration(retryAfter) * time.Second)
			continue
		} else {
			fmt.Println("Error response from server:", res.Status)
			time.Sleep(retryInterval)
			continue
		}
	}

	fmt.Println("Out of retries. Exiting...")
}
