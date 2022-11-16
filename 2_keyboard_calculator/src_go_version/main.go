package main

import (
	"fmt"
	"os"
)

var task string

/* ------------------------------ MAIN FUNCTION ----------------------------- */
// start subroutines for tasks
func main() {
	input_variables()
	while := true
	for while {
		fmt.Println("Write 'start' or 'stop':")
		fmt.Scanln(&task)
		// start or stop program based on user input
		if len(task) > 1 {
			if task == "start" {
				calculator()
			} else if task == "stop" {
				fmt.Println("Program exits.")
				//exit program
				os.Exit(0)
			} else {
				println("Invalid argument. Use 'start' or 'stop'")
			}
		} else {
			println("No argument or invalid argument. Use 'start' or 'stop'")
		}
	}
}
