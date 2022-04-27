package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Power levels criticially low...")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Accessing auxiliary power ... please wait")
	time.Sleep(1000 * time.Millisecond)

	loadingAnimation()
	userInputDispatch()

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

func userInputDispatch() {
	var userInput string = getUserInput()

	scanner := bufio.NewScanner(strings.NewReader(userInput))
	var text string
	if scanner.Scan() {
		text = scanner.Text()
	}
	fmt.Printf("Input was: %q\n", text)

}

func getUserInput() string {

	// fmt.Println()
	// var userInput string
	// fmt.Scanln(&userInput)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")

	return userInput

}
