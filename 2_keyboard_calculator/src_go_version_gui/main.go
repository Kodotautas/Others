package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var task string

/* ------------------------------ MAIN FUNCTION ----------------------------- */
// main function to run the program
func main() {
	a := app.New()
	w := a.NewWindow("Keypress Calculator")
	w.Resize(fyne.NewSize(500, 300))
	// add start and stop buttons
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Choose your task"),
		widget.NewButton("Start", func() {
			task = "start"
		}),
		widget.NewButton("Stop", func() {
			task = "stop"
		}),
	))
	// if stop button is pressed, exit the program
	if task == "stop" {
		w.Close()
	}

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
