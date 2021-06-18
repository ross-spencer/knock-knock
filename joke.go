package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Printing args given knock-knock:", os.Args[1:])

	for _, v := range os.Args[1:] {
		fmt.Println("XXX", v)
		if v == "testtest" {
			fmt.Println("BINGO!")
		}
	}

	if os.Getenv("SUPER_SECRET") == "testtest" {
		fmt.Println("Oh! got it!")
	} else {
		fmt.Println("YYY", os.Getenv("SUPER_SECRET"))
	}

	fmt.Println("Knock knock")
	fmt.Println("Who's there?")
	fmt.Println("Go!")
	fmt.Println("Okay then!")
}
