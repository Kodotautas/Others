package main

import (
	"fmt"
)

var task string

// start subroutines
func main() {
	fmt.Println("Write start or stop:")
	fmt.Scanln(&task)
	if len(task) > 1 {
		if task == "start" {
			calculator()
		} else if task == "stop" {
			println("Stopping...")
		} else {
			println("Invalid argument. Use 'start' or 'stop'")
		}
	} else {
		println("No argument")
	}
}
