package main

import (
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

/* --------------------------------- PROGRAM -------------------------------- */
// main function to run the program
func main() {
	a := app.New()
	w := a.NewWindow("Keypress Calculator")
	w.Resize(fyne.NewSize(450, 300))

	type numericalEntry struct {
		widget.Entry
	}

	interval := numericalEntry{}
	session := numericalEntry{}

	// add start and stop buttons
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Enter interval in seconds:"),
		&interval,
		widget.NewLabel("Enter session in seconds:"),
		&session,
		widget.NewButton("Start", func() {
			calculator(interval, session)
		}),
		widget.NewButton("Exit", func() {
			os.Exit(0)
		}),
	))

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
