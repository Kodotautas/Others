package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var task string

/* ------------------------------ MAIN FUNCTION ----------------------------- */
// main function to run the program
func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
	// fmt.Println("Write 'start' or 'stop':")
	// fmt.Scanln(&task)
	// while := true
	// for while {
	// 	// start or stop program based on user input
	// 	if len(task) > 1 {
	// 		if task == "start" {
	// 			input_variables()
	// 			calculator()
	// 		} else if task == "stop" {
	// 			fmt.Println("Program exits.")
	// 			//exit program
	// 			os.Exit(0)
	// 		} else {
	// 			println("Invalid argument. Use 'start' or 'stop'")
	// 			main()
	// 		}
	// 	} else {
	// 		println("No argument or invalid argument. Use 'start' or 'stop'")
	// 		main()
	// 	}
	// }
}
