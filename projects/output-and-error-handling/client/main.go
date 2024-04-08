package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	res, err := http.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)
	scanner := bufio.NewScanner(res.Body)

	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}