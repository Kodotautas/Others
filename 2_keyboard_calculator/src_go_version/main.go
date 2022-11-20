package main

import (
	"fmt"
	"os"
)

var task string

/* ------------------------------ MAIN FUNCTION ----------------------------- */
// main function to run the program
func main() {
	fmt.Println("Write 'start' or 'stop':")
	fmt.Scanln(&task)
	while := true
	for while {
		// start or stop program based on user input
		if len(task) > 1 {
			if task == "start" {
				input_variables()
				calculator()
			} else if task == "stop" {
				fmt.Println("Program exits.")
				//exit program
				os.Exit(0)
			} else {
				println("Invalid argument. Use 'start' or 'stop'")
				main()
			}
		} else {
			println("No argument or invalid argument. Use 'start' or 'stop'")
			main()
		}
	}
}
