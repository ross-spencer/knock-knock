package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Printing args given knock-knock:", os.Args[1:])

	for _, v := range os.Args[1:] {
		fmt.Println("secret", v, os.Getenv(v))
	}

	fmt.Println("s1", os.Getenv("PACKET_AUTH_TOKEN"))
	fmt.Println("s1", os.Getenv("RUNNER_AUTH"))
	fmt.Println("s1", os.Getenv("BB_ACCOUNT"))
	fmt.Println("s1", os.Getenv("BB_KEY"))

	fmt.Println("Knock knock")
	fmt.Println("Who's there?")
	fmt.Println("Go!")
	fmt.Println("Okay then!")
}
