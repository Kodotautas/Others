package main

import (
	"os"
	"time"
)

var keypresses = []string{}
var keypresses_total = []string{}

// variables
// var interval = 10
// var session = 20
var n int = 3 // number of top keypresses to print

/* ------------------------------ MAIN FUNCTION ----------------------------- */
func calculator() {
	println("")
	println("Started, keypresses will be tracked in", interval, "seconds intervals of", session, "seconds")
	println("")
	println("")
	println("------INTERVAL STATS------")
	for i := 0; i < session; i += interval {
		go store_keypress()
		time.Sleep(time.Duration(interval) * time.Second)
		clearString(keypresses)
		// for key, value := range keypresses {
		// 	println(key, value)
		// }
		println("Keypress in", interval, "seconds:", len(keypresses))
		keypresses_total = append(keypresses_total, keypresses...)
		// calculate typing speed of current session
		top_n_keypresses(n, keypresses, "interval")
		typing_speed(keypresses_total)
		//clear keypresses list
		keypresses = []string{}
		println("--------------------------")
	}
	println("")
	println("------TOTAL STATS------")
	println("Keypress in", session, "seconds:", len(keypresses_total))
	// top n keypresses_total of list
	top_n_keypresses(n, keypresses_total, "session")
	// calculate typing speed of current session
	typing_speed(keypresses_total)
	println("--------------------------")
	os.Exit(0)
}
