package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("PORT: ", portString)
}
