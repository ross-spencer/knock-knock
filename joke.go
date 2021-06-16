package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Printing args given knock-knock:", os.Args[1:])

	fmt.Println("Knock knock")
	fmt.Println("Who's there?")
	fmt.Println("Go")
	fmt.Println("Okay then!")
}
