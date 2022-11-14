package main

import (
	"time"
)

var keypresses = []string{}
var keypresses_total = []string{}

// integer variables
var interval = 10
var session = 20
var n int = 3 // number of top keypresses to print

// // store each keypresses separetely as number of keypresses
// func store_keypress() {
// 	println("")
// 	println("Waiting for keypresses...")
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		char, _, _ := reader.ReadRune()
// 		keypresses = append(keypresses, string(char))
// 	}
// }

// // remove non alphanumeric characters from keypresses list
// var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

// func clearString(keypresses []string) {
// 	for i, v := range keypresses {
// 		keypresses[i] = nonAlphanumericRegex.ReplaceAllString(v, "")
// 	}
// }

// // top n keypresses of list
// func top_n_keypresses(n int, keypresses []string) {
// 	// sort keypresses
// 	sort.Strings(keypresses)
// 	// count keypresses
// 	count := make(map[string]int)
// 	for _, v := range keypresses {
// 		count[v]++
// 	}
// 	// sort keypresses by count
// 	type kv struct {
// 		Key   string
// 		Value int
// 	}
// 	var ss []kv
// 	for k, v := range count {
// 		ss = append(ss, kv{k, v})
// 	}
// 	sort.Slice(ss, func(i, j int) bool {
// 		return ss[i].Value > ss[j].Value
// 	})
// 	// print top n keypresses
// 	if len(ss) < n {
// 		n = len(ss)
// 	}
// 	println("Top", n, "keypresses:")
// 	for i := 0; i < n; i++ {
// 		fmt.Printf("%s: %d, ", ss[i].Key, ss[i].Value)
// 	}
// }

// // calculate typing speed of current session
// func typing_speed(list []string) {
// 	// calculate typing speed
// 	var speed float64 = float64(len(list)) / (float64(session) / 60)
// 	// print typing speed
// 	println("")
// 	fmt.Printf("Speed of session: %.1f keypresses per minute", speed)
// 	fmt.Println()
// }

/* ------------------------------ MAIN FUNCTION ----------------------------- */
func main() {
	println("")
	println("Started, keypresses will be tracked in", interval, "seconds intervals of", session, "seconds")
	println("")
	//for every interval of time print number of keypresses
	for i := 0; i < session; i += interval {
		go store_keypress()
		time.Sleep(time.Duration(interval) * time.Second)
		println("")
		println("------INTERVAL STATS------")
		clearString(keypresses)
		// for key, value := range keypresses {
		// 	println(key, value)
		// }
		println("Keypress in", interval, "seconds:", len(keypresses))
		keypresses_total = append(keypresses_total, keypresses...)
		// calculate typing speed of current session
		top_n_keypresses(n, keypresses)
		typing_speed(keypresses_total)
		//clear keypresses list
		keypresses = []string{}
		println("--------------------------")
	}
	println("")
	println("------TOTAL STATS------")
	println("Keypress in", session, "seconds:", len(keypresses_total))
	// top n keypresses_total of list
	top_n_keypresses(n, keypresses_total)
	// calculate typing speed of current session
	typing_speed(keypresses_total)
}
