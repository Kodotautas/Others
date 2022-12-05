package main

import (
	"log"
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

	interval := widget.NewEntry()
	interval.SetPlaceHolder("Enter interval in seconds:")

	session := widget.NewEntry()
	session.SetPlaceHolder("Enter session in seconds:")

	// add start and stop buttons
	w.SetContent(widget.NewVBox(
		interval,
		session,
		widget.NewButton("Start", func() {
			log.Println(interval.Text)
		}),
		widget.NewButton("Exit app", func() {
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
