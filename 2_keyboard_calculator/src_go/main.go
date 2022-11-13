package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
	"time"
)

// func variables() {
// 	interval := flag.Float64("interval", 1.0, "interval in seconds")
// 	session := flag.Float64("session", 60.0, "session in seconds")
// 	n := flag.Int("n", 5, "top n keypresses")
// 	flag.Parse()
// }

var keypresses map[interface{}]interface{}
var keypress_total_dict map[interface{}]interface{}
var keypresses_total []interface{}
var keypress_speed []interface{}

var interval float64
var session float64

// store keypresses in keypress dictionary
func store_keypress(keypress interface{}) {
	if keypresses[keypress] == nil {
		keypresses[keypress] = 1
	} else {
		keypresses[keypress] = keypresses[keypress].(int) + 1
	}
}

// store total keypresses in dictionary
func store_keypress_total(keypress interface{}) {
	if keypress_total_dict[keypress] == nil {
		keypress_total_dict[keypress] = 1
	} else {
		keypress_total_dict[keypress] = keypress_total_dict[keypress].(int) + 1
	}
}

// sum how many keypresses in interval and store in dictionary
func count_keypresses() {
	for i := 0; i < len(keypresses_total); i++ {
		store_keypress(keypresses_total[i])
	}
}

// return top n keypresses from store_keypresses dictionary
func top_n_keypresses(n int) []interface{} {
	var top_n []interface{}
	for i := 0; i < n; i++ {
		var max_key interface{}
		var max_value int
		for key, value := range keypresses {
			if value.(int) > max_value {
				max_value = value.(int)
				max_key = key
			}
		}
		top_n = append(top_n, max_key)
		delete(keypresses, max_key)
	}
	return top_n
}

// calculate average keypresses speed per minute
func calculate_speed() float64 {
	var sum int
	for _, value := range keypress_speed {
		sum += value.(int)
	}
	return float64(sum) / 60 //TO DO: make it dynamic
}
func main() {
	fmt.Println(func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("Started, keypresses will be tracked every {{.interval}} seconds for {{.session}} seconds.")).Execute(&buf, map[string]interface {
		}{"interval": interval, "session": session})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}())
	fmt.Println("")
	fmt.Println("------------INTERVAL STATS----------")
	start_time := float64(time.Now().UnixNano()) / 1000000000.0
	for float64(time.Now().UnixNano())/1000000000.0-start_time < session {
		fmt.Println(func() string {
			var buf bytes.Buffer
			err := template.Must(template.New("f").Parse("Keypress in {{.interval}} seconds: ")).Execute(&buf, map[string]interface {
			}{"interval": interval})
			if err != nil {
				panic(err)
			}
			return buf.String()
		}(), count_keypresses(keypresses))
		fmt.Println(func() string {
			var buf bytes.Buffer
			err := template.Must(template.New("f").Parse("Top 3 keypresses: ")).Execute(&buf, map[string]interface {
			}{})
			if err != nil {
				panic(err)
			}
			return buf.String()
		}(), top_n_keypresses(keypresses, n))
		fmt.Println(func() string {
			var buf bytes.Buffer
			err := template.Must(template.New("f").Parse("Current session speed:")).Execute(&buf, map[string]interface {
			}{})
			if err != nil {
				panic(err)
			}
			return buf.String()
		}(), calculate_average_speed(keypress_speed), "characters per minute")
		store_total_keypresses(keypresses)
		keypresses := nil
		fmt.Println("------------------------------------")
	}
	fmt.Println("")
	fmt.Println("------------SESSION STATS----------")
	fmt.Println(func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("Total keypresses: ")).Execute(&buf, map[string]interface {
		}{})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}(), sum(keypresses_total))
	fmt.Println(func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("Top 3 keypresses: ")).Execute(&buf, map[string]interface {
		}{})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}(), top_n_keypresses(keypress_total_dict, n))
	fmt.Println(func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("Final session speed:")).Execute(&buf, map[string]interface {
		}{})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}(), calculate_average_speed(keypress_speed), "characters per minute")
	fmt.Println("")
}
func init() {
	if session < interval {
		fmt.Println("Session time must be greater than interval time. Check it and try again.")
		os.Exit(0)
	} else {
		main()
	}
}
