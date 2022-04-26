package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Power levels criticially low...")
	time.Sleep(250 * time.Millisecond)
	fmt.Println("Accessing auxiliary power ... please wait")
	time.Sleep(250 * time.Millisecond)

	loadingAnimation()
	getUserInput()

}

func loadingAnimation() {

	dots := "."
	var count int = 0

	for count < 20 {

		fmt.Print(dots)
		count += 1
		time.Sleep(250 * time.Millisecond)

	}

}

func getUserInput() {

	fmt.Println()
	fmt.Println("Ready for input: ")
	var userInput string
	fmt.Scanln(&userInput)
	fmt.Print("Your input: ")
	fmt.Print(userInput)

}

func help() {

}
