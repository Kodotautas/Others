// program import helper.go file and use it
package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
	"time"
)

func main() {
	keys_tracker := Keys_Tracker()
	fmt.Println(func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("Started, keypresses will be tracked every {{.interval}} seconds for {{.session}} seconds.")).
			Execute(&buf, map[string]interface{}{"interval": interval, "session": session})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}())
	fmt.Println("")
	fmt.Println("------------INTERVAL STATS----------")
	for float64(time.Now().UnixNano())/1000000000.0-keys_tracker.start_time < keys_tracker.session {
		fmt.Println(func() string {
			var buf bytes.Buffer
			err := template.Must(template.New("f").Parse("Keypress in {{.interval}} seconds: ")).
				Execute(&buf, map[string]interface{}{"interval": interval})
			if err != nil {
				panic(err)
			}
			return buf.String()
		}(), keys_tracker.count_keypresses(keypresses))
		fmt.Println(func() string {
			var buf bytes.Buffer
			err := template.Must(template.New("f").Parse("Top 3 keypresses: ")).
				Execute(&buf, map[string]interface{}{})
			if err != nil {
				panic(err)
			}
			return buf.String()
		}(), keys_tracker.top_n_keypresses(keypresses, n))
		fmt.Println(func() string {
			var buf bytes.Buffer
			err := template.Must(template.New("f").Parse("Current session speed:")).
				Execute(&buf, map[string]interface{}{})
			if err != nil {
				panic(err)
			}
			return buf.String()
		}(), keys_tracker.calculate_average_speed(keypress_speed), "characters per minute")
		keys_tracker.store_total_keypresses(keypresses)
		keypresses := nil
		fmt.Println("------------------------------------")
	}
	fmt.Println("")
	fmt.Println("------------SESSION STATS----------")
	fmt.Println(func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("Total keypresses: ")).
			Execute(&buf, map[string]interface{}{})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}(), sum(keypresses_total))
	fmt.Println(func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("Top 3 keypresses: ")).
			Execute(&buf, map[string]interface{}{})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}(), keys_tracker.top_n_keypresses(keypress_total_dict, n))
	fmt.Println(func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("Final session speed:")).
			Execute(&buf, map[string]interface{}{})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}(), keys_tracker.calculate_average_speed(keypress_speed), "characters per minute")
	fmt.Println("")
}

func Keys_Tracker() {
	panic("unimplemented")
}

func init() {
	if session < interval {
		fmt.Println("Session time must be greater than interval time. Check it and try again.")
		os.Exit(0)
	} else {
		main()
	}
}
