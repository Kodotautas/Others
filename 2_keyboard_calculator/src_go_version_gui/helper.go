package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

var interval int
var session int

// get interval and session from user
func input_variables() {
	// //interval must be greater than 0 and not string
	fmt.Println("Enter interval in seconds:")
	fmt.Scanln(&interval)
	if interval <= 0 || interval == 0 {
		fmt.Println("Interval must be greater than 0 and not string.")
	}
	// session variable must be greater than 0 and not string
	fmt.Println("Enter session in seconds:")
	fmt.Scanln(&session)
	if session <= 0 || session == 0 {
		fmt.Println("Session must be greater than 0 and not string.")
	}
	// session varianle must be equal or greater than interval variable
	if session < interval {
		fmt.Println("Session must be equal or greater than interval. Try again.")
	}
}

// store each keypresses separetely as number of keypresses
func store_keypress() {
	println("")
	println("Waiting for keypresses... and press enter in the end")
	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, _ := reader.ReadRune()
		keypresses = append(keypresses, string(char))
	}
}

// remove special characters from string
func removeSpecialChars(list []string) {
	for i := 0; i < len(list); i++ {
		re := regexp.MustCompile("[^a-zA-Z0-9]+")
		list[i] = re.ReplaceAllString(list[i], "Special_char")
	}
}

// top n keypresses of list
func top_n_keypresses(n int, keypresses []string, period string) {
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
	println("Top", n, period, "keypresses:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s: %d, ", ss[i].Key, ss[i].Value)
	}
}

// calculate typing speed of current session
func typing_speed(list []string) {
	// typing speed calculation
	var speed float64 = float64(len(list)) / (float64(session) / 60)
	// print typing speed
	println("")
	fmt.Printf("Speed of session: %.1f keypresses per minute", speed)
	fmt.Println()
}

// convert text to integer
func convertToInt(text string) int {
	var number int
	fmt.Sscan(text, &number)
	return number
}
