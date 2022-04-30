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
	// var userInput string = getUserInput()

	var exit bool = false

	for exit != true {

		userInput := getUserInput()
		if userInput[0] == "help" {
			printHelpMenu()
			userInput = getUserInput()
		} else if userInput[0] == "exit" {
			exit = true
		}
	}

	// scanner := bufio.NewScanner(strings.NewReader(userInput))
	// scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

}

func getUserInput() []string {

	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	userInput, _ := reader.ReadString('\n')
	userInputArray := strings.Fields(userInput)
	return userInputArray

	// fmt.Scanln(&userInput)
	// nameThing := strings.NewReader(name)

	// fmt.Printf("%q", nameArray[1])

	// reader := bufio.NewReader(os.Stdin)
	// reader := bufio.NewReader(userInput)
	// scanner := bufio.NewScanner(strings.NewReader(userInput));
	// scanner.Split(bufio.ScanWords);
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	// fmt.Print("> ")

	// return userInput

}

func printHelpMenu() {
	fmt.Println("Commands	-	Description")
	fmt.Println("archive 	- 	Accesses archives")
}
