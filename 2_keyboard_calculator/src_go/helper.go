package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

// store each keypresses separetely as number of keypresses
func store_keypress() {
	println("")
	println("Waiting for keypresses...")
	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, _ := reader.ReadRune()
		keypresses = append(keypresses, string(char))
	}
}

// func store_keypress() {
// 	println("")
// 	println("Waiting for keypresses...")
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		char, _, _ := reader.ReadRune()
// 		keypresses :=
// 		keypresses = append(keypresses, string(char))
// 	}
// }

// remove non alphanumeric characters from keypresses list
var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func clearString(keypresses []string) {
	for i, v := range keypresses {
		keypresses[i] = nonAlphanumericRegex.ReplaceAllString(v, "")
	}
}

// top n keypresses of list
func top_n_keypresses(n int, keypresses []string) {
	// sort keypresses
	sort.Strings(keypresses)
	// count keypresses
	count := make(map[string]int)
	for _, v := range keypresses {
		count[v]++
	}
	// sort keypresses by count
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range count {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	// print top n keypresses
	if len(ss) < n {
		n = len(ss)
	}
	println("Top", n, "keypresses:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s: %d, ", ss[i].Key, ss[i].Value)
	}
}

// calculate typing speed of current session
func typing_speed(list []string) {
	// calculate typing speed
	var speed float64 = float64(len(list)) / (float64(session) / 60)
	// print typing speed
	println("")
	fmt.Printf("Speed of session: %.1f keypresses per minute", speed)
	fmt.Println()
}
