package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:8080")
	if err != nil {

		if res != nil {
			fmt.Println("Response status:", res.Status)
			if res.StatusCode == http.StatusTooManyRequests {
				fmt.Println("Server is too busy. Retry after:", res.Header.Get("Retry-After"))
			} else {
				fmt.Println("Failed to get response from server:", err)
			}
		} else {
			fmt.Println("Failed to connect to server:", err)
		}
		return
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
	} else {
		fmt.Println("Error response from server:", res.Status)
	}
}
