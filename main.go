package main

import (
	"fmt"
	"os"
)

func main() {
	nickname := os.Getenv("NICKNAME")

	if nickname == "" {
		nickname = "John Doe"
	}

	fmt.Printf("Hello, %s!\n", nickname)
}